package suffiks

import (
	"encoding/json"

	"github.com/suffiks/suffiks-tinygo/internal/env"
	"github.com/suffiks/suffiks-tinygo/protogen"
)

func AddEnv(name, value string) {
	v := &protogen.KeyValue{
		Name:  name,
		Value: value,
	}

	env.AddLabel(byteToPtr(marshal(v)))
}

func AddEnvFrom(name string, typ protogen.EnvFromType, optional bool) {
	v := &protogen.EnvFrom{
		Name:     name,
		Optional: optional,
		Type:     typ,
	}

	env.AddEnvFrom(byteToPtr(marshal(v)))
}

func AddLabel(name, value string) {
	v := &protogen.KeyValue{
		Name:  name,
		Value: value,
	}

	env.AddLabel(byteToPtr(marshal(v)))
}

func AddAnnotation(name, value string) {
	v := &protogen.KeyValue{
		Name:  name,
		Value: value,
	}

	env.AddAnnotation(byteToPtr(marshal(v)))
}

func AddInitContainer(container *protogen.Container) {
	panic("not implemented")
}

func AddSidecar(container *protogen.Container) {
	panic("not implemented")
}

func MergePatch(patch []byte) {
	env.MergePatch(byteToPtr(patch))
}

func ValidationError(path, detail, value string) {
	message := &protogen.ValidationError{
		Path:   path,
		Detail: detail,
		Value:  value,
	}

	env.ValidationError(byteToPtr(marshal(message)))
}

func GetOwner() *protogen.Owner {
	ptrAndSize := env.GetOwner()
	owner := &protogen.Owner{}

	b := getByteSlice(ptrAndSize)
	if err := owner.UnmarshalVT(b); err != nil {
		panic(err)
	}

	return owner
}

func GetSpec[T any](v T) T {
	b := getByteSlice(env.GetSpec())
	if err := json.Unmarshal(b, &v); err != nil {
		panic(err)
	}

	return v
}

func DefaultingResponse(v any) uint64 {
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}

	return byteToPtrAndSize(b)
}

func GetResource[T any](group, version, resource, name string) (T, error) {
	gvr := &protogen.GroupVersionResource{
		Group:    ptrString(group),
		Version:  ptrString(version),
		Resource: ptrString(resource),
	}

	gvrPtr, gvrSize := byteToPtr(marshal(gvr))
	namePtr, nameSize := stringToWasiPtr(name)

	ptrAndSize := env.GetResource(gvrPtr, gvrSize, namePtr, nameSize)
	if ptrAndSize <= 20 {
		return *new(T), ClientError{Code: ptrAndSize}
	}

	b := getByteSlice(ptrAndSize)
	var v T
	if err := json.Unmarshal(b, &v); err != nil {
		return *new(T), err
	}

	return v, nil
}

func DeleteResource(group, version, resource, name string) error {
	gvr := &protogen.GroupVersionResource{
		Group:    ptrString(group),
		Version:  ptrString(version),
		Resource: ptrString(resource),
	}

	gvrPtr, gvrSize := byteToPtr(marshal(gvr))
	namePtr, nameSize := stringToWasiPtr(name)

	ptrAndSize := env.DeleteResource(gvrPtr, gvrSize, namePtr, nameSize)
	if ptrAndSize > 0 {
		return ClientError{Code: ptrAndSize}
	}

	return nil
}

func CreateResource[T any](group, version, resource string, spec any) (T, error) {
	gvr := &protogen.GroupVersionResource{
		Group:    ptrString(group),
		Version:  ptrString(version),
		Resource: ptrString(resource),
	}

	specBytes, err := json.Marshal(spec)
	if err != nil {
		return *new(T), err
	}

	gvrPtr, gvrSize := byteToPtr(marshal(gvr))
	specPtr, specSize := byteToPtr(specBytes)

	ptrAndSize := env.CreateResource(gvrPtr, gvrSize, specPtr, specSize)
	if ptrAndSize <= 20 {
		return *new(T), ClientError{Code: ptrAndSize}
	}

	b := getByteSlice(ptrAndSize)
	var v T
	if err := json.Unmarshal(b, &v); err != nil {
		return *new(T), err
	}

	return v, nil
}

func UpdateResource[T any](group, version, resource string, spec any) (T, error) {
	gvr := &protogen.GroupVersionResource{
		Group:    ptrString(group),
		Version:  ptrString(version),
		Resource: ptrString(resource),
	}

	specBytes, err := json.Marshal(spec)
	if err != nil {
		return *new(T), err
	}

	gvrPtr, gvrSize := byteToPtr(marshal(gvr))
	specPtr, specSize := byteToPtr(specBytes)

	ptrAndSize := env.UpdateResource(gvrPtr, gvrSize, specPtr, specSize)
	if ptrAndSize <= 20 {
		return *new(T), ClientError{Code: ptrAndSize}
	}

	b := getByteSlice(ptrAndSize)
	var v T
	if err := json.Unmarshal(b, &v); err != nil {
		return *new(T), err
	}

	return v, nil
}
