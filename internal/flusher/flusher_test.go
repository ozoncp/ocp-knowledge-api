package flusher_test

import (
	"errors"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/ozoncp/ocp-knowledge-api/internal/flusher"
	"github.com/ozoncp/ocp-knowledge-api/internal/mocks"
	"github.com/ozoncp/ocp-knowledge-api/internal/models"
)

var _ = Describe("Flusher", func() {
	var (
		mockRepo *mocks.MockKnowledgeRepo
		mockCtrl *gomock.Controller
		fl       flusher.Flusher
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockRepo = mocks.NewMockKnowledgeRepo(mockCtrl)
		fl = flusher.NewFlusher(2, mockRepo)
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Context("Flusher - Happy paths", func() {
		It("Flush 3 items - Success", func() {
			mockRepo.EXPECT().
				AddKnowledges(gomock.Any()).
				Times(2).
				Return(nil)

			res := fl.Flush([]models.Knowledge{
				{
					Id:     1,
					UserId: 1,
					Topic:  1,
					Text:   "123",
				},
				{
					Id:     2,
					UserId: 2,
					Topic:  2,
					Text:   "123",
				},
				{
					Id:     3,
					UserId: 3,
					Topic:  3,
					Text:   "123",
				},
			})

			Expect(res).Should(BeNil())
		})

		It("Flush 1 items - Success", func() {
			mockRepo.EXPECT().
				AddKnowledges(gomock.Any()).
				Times(1).
				Return(nil)

			res := fl.Flush([]models.Knowledge{
				{
					Id:     1,
					UserId: 1,
					Topic:  1,
					Text:   "123",
				},
			})

			Expect(res).Should(BeNil())
		})
	})

	Context("Flusher - Errors", func() {
		It("Flush 0 items - Expecting error", func() {
			res := fl.Flush([]models.Knowledge{})
			Expect(res).Should(BeEquivalentTo([]models.Knowledge{}))
		})

		It("Error occured during adding knowledge", func() {
			mockRepo.EXPECT().
				AddKnowledges(gomock.Any()).
				Times(1).
				Return(errors.New("error"))

			entities := []models.Knowledge{
				{
					Id:     1,
					UserId: 1,
					Topic:  1,
					Text:   "123",
				},
			}
			res := fl.Flush(entities)

			Expect(res).Should(BeEquivalentTo(entities))
		})
	})
})
