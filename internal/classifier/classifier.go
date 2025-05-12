package classifier

import (
	"context"
	"fmt"

	"github.com/nlpodyssey/cybertron/pkg/tasks"
	"github.com/nlpodyssey/cybertron/pkg/tasks/zeroshotclassifier"
)

// ZeroShot is a struct that holds the Zero-Shot Classifier model.
type ZeroShot struct {
	model zeroshotclassifier.Interface
}

// NewZeroShot initializes a new ZeroShot classifier with the specified model directory and model name.
func NewZeroShot(modelsDir, model string) (*ZeroShot, error) {
	m, err := tasks.Load[zeroshotclassifier.Interface](&tasks.Config{ModelsDir: modelsDir, ModelName: model})
	if err != nil {
		return nil, fmt.Errorf("failed to load model: %w", err)
	}

	return &ZeroShot{model: m}, nil
}

// Classify performs zero-shot classification on the provided text using the specified labels.
func (z *ZeroShot) Classify(ctx context.Context, text string, labels []string) (*zeroshotclassifier.Response, error) {
	params := zeroshotclassifier.Parameters{
		CandidateLabels:    labels,
		HypothesisTemplate: zeroshotclassifier.DefaultHypothesisTemplate,
		MultiLabel:         true,
	}

	result, err := z.model.Classify(ctx, text, params)
	if err != nil {
		return nil, fmt.Errorf("failed to classify text: %w", err)
	}

	return &result, nil
}
