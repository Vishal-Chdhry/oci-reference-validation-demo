package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/go-containerregistry/pkg/crane"
	"github.com/google/go-containerregistry/pkg/name"
	"github.com/google/go-containerregistry/pkg/v1/remote"
	notationregistry "github.com/notaryproject/notation-go/registry"
	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
)

type repositoryClient struct {
	ref        name.Reference
	craneOpts  crane.Option
	remoteOpts []remote.Option
}

func NewRepository(craneOpts crane.Option, remoteOpts []remote.Option, ref name.Reference) notationregistry.Repository {
	return &repositoryClient{
		craneOpts:  craneOpts,
		remoteOpts: remoteOpts,
		ref:        ref,
	}
}

func (c *repositoryClient) Resolve(ctx context.Context, reference string) (ocispec.Descriptor, error) {
	head, err := crane.Head(reference)
	if err != nil {
		return ocispec.Descriptor{}, nil
	}
	descriptor := v1ToOciSpecDescriptor(*head)
	return descriptor, nil
}

func (c *repositoryClient) ListSignatures(ctx context.Context, desc ocispec.Descriptor, fn func(signatureManifests []ocispec.Descriptor) error) error {
	referrers, err := remote.Referrers(c.ref.Context().Digest(desc.Digest.String()), c.remoteOpts...)
	if err != nil {
		return err
	}

	referrersDescs, err := referrers.IndexManifest()
	if err != nil {
		return err
	}

	descList := []ocispec.Descriptor{}
	for _, d := range referrersDescs.Manifests {
		if d.ArtifactType == notationregistry.ArtifactTypeNotation {
			descList = append(descList, v1ToOciSpecDescriptor(d))
		}
	}

	return fn(descList)
}

func (c *repositoryClient) FetchSignatureBlob(ctx context.Context, desc ocispec.Descriptor) ([]byte, ocispec.Descriptor, error) {
	manifestRef := c.getReferenceFromDescriptor(desc)

	manifestBytes, err := crane.Manifest(manifestRef)
	if err != nil {
		return nil, ocispec.Descriptor{}, err
	}

	var manifest ocispec.Manifest
	if err := json.Unmarshal(manifestBytes, &manifest); err != nil {
		return nil, ocispec.Descriptor{}, err
	}
	manifestDesc := manifest.Layers[0]

	signatureBlobRef := c.getReferenceFromDescriptor(manifestDesc)

	signatureBlobLayer, err := crane.PullLayer(signatureBlobRef)
	if err != nil {
		panic(err)
	}

	io, err := signatureBlobLayer.Uncompressed()
	if err != nil {
		panic(err)
	}
	SigBlobBuf := new(bytes.Buffer)

	_, err = SigBlobBuf.ReadFrom(io)
	if err != nil {
		panic(err)
	}
	return SigBlobBuf.Bytes(), manifestDesc, nil
}

func (c *repositoryClient) PushSignature(ctx context.Context, mediaType string, blob []byte, subject ocispec.Descriptor, annotations map[string]string) (blobDesc, manifestDesc ocispec.Descriptor, err error) {
	return ocispec.Descriptor{}, ocispec.Descriptor{}, fmt.Errorf("push signature is not implemented")
}