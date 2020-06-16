package sequence

import (
	"bytes"
	"context"
	"encoding/json"
	"image"
	"image/jpeg"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/nfnt/resize"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"

	"github.com/aubm/random-go-tips/pkg/config"
	"github.com/aubm/random-go-tips/pkg/webserver"
)

type Message struct {
	Ctx  context.Context
	Data []byte
}

func Run(config config.Config) {
	cfg := jaegercfg.Configuration{
		ServiceName: "imageResizer",
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans: true,
		},
	}
	tracer, closer, err := cfg.NewTracer()
	if err != nil {
		log.Fatalf("could not initialize jaeger tracer: %v", err.Error())
	}
	opentracing.SetGlobalTracer(tracer)
	defer closer.Close()

	urlChan := make(chan Message)
	imageChan := make(chan Message)
	resizedDataChan := make(chan Message)

	go downloadImages(urlChan, imageChan)
	go resizeImages(imageChan, resizedDataChan)
	go readResizedImages(resizedDataChan)

	webserver.Start(config.WebAppAddr, func(w http.ResponseWriter, r *http.Request) {
		span, ctx := opentracing.StartSpanFromContext(r.Context(), "newImagesResizeRequest")
		defer span.Finish()

		var imagesUrl []string
		if err := json.NewDecoder(r.Body).Decode(&imagesUrl); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		for _, imageUrl := range imagesUrl {
			_, ctx = opentracing.StartSpanFromContext(ctx, "handleImage", opentracing.ChildOf(span.Context()))
			urlChan <- Message{
				Ctx:  ctx,
				Data: []byte(imageUrl),
			}
		}
	})
}

func downloadImages(in chan Message, out chan Message) {
	for {
		go func(msg Message) {
			span := opentracing.SpanFromContext(msg.Ctx)
			childSpan := span.Tracer().StartSpan("downloadImages", opentracing.ChildOf(span.Context()))
			defer childSpan.Finish()

			url := string(msg.Data)
			resp, err := http.Get(url)
			if err != nil {
				log.Printf("failed to download %v: %v", url, err)
				return
			}
			defer resp.Body.Close()

			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Printf("failed to read response body for %v: %v", url, err)
				return
			}

			out <- Message{
				Ctx:  msg.Ctx,
				Data: b,
			}
		}(<-in)
	}
}

func resizeImages(in chan Message, out chan Message) {
	for {
		go func(msg Message) {
			span := opentracing.SpanFromContext(msg.Ctx)
			childSpan := span.Tracer().StartSpan("resizeImages", opentracing.ChildOf(span.Context()))
			defer childSpan.Finish()

			img, _, err := image.Decode(bytes.NewBuffer(msg.Data))
			if err != nil {
				log.Printf("failed to decode image from data: %v", err)
				return
			}

			newImage := resize.Resize(50, 0, img, resize.Lanczos3)
			b := new(bytes.Buffer)
			if err := jpeg.Encode(b, newImage, nil); err != nil {
				log.Printf("failed to encode resized image: %v", err)
				return
			}

			out <- Message{
				Ctx:  msg.Ctx,
				Data: b.Bytes(),
			}
		}(<-in)
	}
}

func readResizedImages(msgs chan Message) {
	for {
		go func(msg Message) {
			span := opentracing.SpanFromContext(msg.Ctx)
			defer span.Finish()

			childSpan := span.Tracer().StartSpan("readResizedImages", opentracing.ChildOf(span.Context()))
			defer childSpan.Finish()

			time.Sleep(time.Second)

			log.Print("read resized image data")
		}(<-msgs)
	}
}
