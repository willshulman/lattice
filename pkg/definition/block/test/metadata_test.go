package test

import (
	"reflect"
	"testing"

	"github.com/mlab-lattice/system/pkg/definition/block"
)

func TestMetadata_Validate(t *testing.T) {
	Validate(
		t,
		nil,

		// Invalid Metadata
		[]ValidateTest{
			{
				Description:     "empty",
				DefinitionBlock: &block.Metadata{},
			},
			{
				Description: "no Type",
				DefinitionBlock: &block.Metadata{
					Name: "my-system",
				},
			},
			{
				Description: "no Name",
				DefinitionBlock: &block.Metadata{
					Type: "my-type",
				},
			},
		},

		// Valid Metadata
		[]ValidateTest{
			{
				Description: "Name and Type",
				DefinitionBlock: &block.Metadata{
					Name: "my-system",
					Type: "my-type",
				},
			},
			{
				Description: "Name, Type, and Description",
				DefinitionBlock: &block.Metadata{
					Name:        "my-system",
					Type:        "my-type",
					Description: "this is my system",
				},
			},
		},
	)
}

func TestMetadata_JSON(t *testing.T) {
	JSON(
		t,
		reflect.TypeOf(block.Metadata{}),
		[]JSONTest{
			{
				Description: "MockSystemMetadata",
				Bytes:       MockSystemMetadataExpectedJSON(),
				ValuePtr:    MockSystemMetadata(),
			},
			{
				Description: "MockServiceMetadata",
				Bytes:       MockServiceMetadataExpectedJSON(),
				ValuePtr:    MockServiceMetadata(),
			},
			{
				Description: "MockServiceDifferentNameMetadata",
				Bytes:       MockServiceDifferentNameMetadataExpectedJSON(),
				ValuePtr:    MockServiceDifferentNameMetadata(),
			},
		},
	)
}

// TODO: add MetadataParameter
