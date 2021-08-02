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
		mockRepo *mocks.MockRepo
		mockCtrl *gomock.Controller
		fl       flusher.Flusher
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockRepo = mocks.NewMockRepo(mockCtrl)
		fl = flusher.NewFlusher(2, mockRepo)
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Context("Flusher - Happy paths", func() {
		It("Flush 3 items - Success", func() {
			mockRepo.EXPECT().
				AddKnowledge(gomock.Any()).
				Times(2).
				Return(nil)

			res, err := fl.Flush([]models.Knowledge{
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
			Expect(err).ShouldNot(HaveOccurred())
		})

		It("Flush 1 items - Success", func() {
			mockRepo.EXPECT().
				AddKnowledge(gomock.Any()).
				Times(1).
				Return(nil)

			res, err := fl.Flush([]models.Knowledge{
				{
					Id:     1,
					UserId: 1,
					Topic:  1,
					Text:   "123",
				},
			})

			Expect(res).Should(BeNil())
			Expect(err).ShouldNot(HaveOccurred())
		})
	})

	Context("Flusher - Errors", func() {
		It("Flush 0 items - Expecting error", func() {
			res, err := fl.Flush([]models.Knowledge{})

			Expect(res).Should(BeNil())
			Expect(err).Should(HaveOccurred())
		})

		It("Error occured during adding knowledge", func() {
			mockRepo.EXPECT().
				AddKnowledge(gomock.Any()).
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
			res, err := fl.Flush(entities)

			Expect(res).Should(BeEquivalentTo(entities))
			Expect(err).Should(HaveOccurred())
		})
	})
})
