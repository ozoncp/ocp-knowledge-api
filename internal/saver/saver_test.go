package saver_test

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	"github.com/ozoncp/ocp-knowledge-api/internal/mocks"
	"github.com/ozoncp/ocp-knowledge-api/internal/models"
	"github.com/ozoncp/ocp-knowledge-api/internal/saver"
	"time"
)

var _ = Describe("Saver", func() {
	var (
		ctrl    *gomock.Controller
		flusher *mocks.MockFlusher
		data    []*models.Knowledge
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		flusher = mocks.NewMockFlusher(ctrl)
		data = []*models.Knowledge{
			{Id: 1, UserId: 2, Topic: 3, Text: "123"},
			{Id: 4, UserId: 5, Topic: 6, Text: "456"},
			{Id: 7, UserId: 8, Topic: 9, Text: "789"},
		}
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Context("Saver", func() {
		It("Flushes knowledge by timeout", func() {
			flusher.EXPECT().Flush(gomock.Any()).MinTimes(2)
			saver := saver.NewSaver(3, flusher)

			for _, v := range data {
				saver.Save(*v)
			}

			time.Sleep(time.Second * 1)
			saver.Close()
		})
	})
})
