package fake

import (
	lattice_v1 "github.com/mlab-lattice/system/pkg/backend/kubernetes/customresource/apis/lattice/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeTeardowns implements TeardownInterface
type FakeTeardowns struct {
	Fake *FakeLatticeV1
	ns   string
}

var teardownsResource = schema.GroupVersionResource{Group: "lattice.mlab.com", Version: "v1", Resource: "teardowns"}

var teardownsKind = schema.GroupVersionKind{Group: "lattice.mlab.com", Version: "v1", Kind: "Teardown"}

// Get takes name of the teardown, and returns the corresponding teardown object, and an error if there is any.
func (c *FakeTeardowns) Get(name string, options v1.GetOptions) (result *lattice_v1.Teardown, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(teardownsResource, c.ns, name), &lattice_v1.Teardown{})

	if obj == nil {
		return nil, err
	}
	return obj.(*lattice_v1.Teardown), err
}

// List takes label and field selectors, and returns the list of Teardowns that match those selectors.
func (c *FakeTeardowns) List(opts v1.ListOptions) (result *lattice_v1.TeardownList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(teardownsResource, teardownsKind, c.ns, opts), &lattice_v1.TeardownList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &lattice_v1.TeardownList{}
	for _, item := range obj.(*lattice_v1.TeardownList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested teardowns.
func (c *FakeTeardowns) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(teardownsResource, c.ns, opts))

}

// Create takes the representation of a teardown and creates it.  Returns the server's representation of the teardown, and an error, if there is any.
func (c *FakeTeardowns) Create(teardown *lattice_v1.Teardown) (result *lattice_v1.Teardown, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(teardownsResource, c.ns, teardown), &lattice_v1.Teardown{})

	if obj == nil {
		return nil, err
	}
	return obj.(*lattice_v1.Teardown), err
}

// Update takes the representation of a teardown and updates it. Returns the server's representation of the teardown, and an error, if there is any.
func (c *FakeTeardowns) Update(teardown *lattice_v1.Teardown) (result *lattice_v1.Teardown, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(teardownsResource, c.ns, teardown), &lattice_v1.Teardown{})

	if obj == nil {
		return nil, err
	}
	return obj.(*lattice_v1.Teardown), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeTeardowns) UpdateStatus(teardown *lattice_v1.Teardown) (*lattice_v1.Teardown, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(teardownsResource, "status", c.ns, teardown), &lattice_v1.Teardown{})

	if obj == nil {
		return nil, err
	}
	return obj.(*lattice_v1.Teardown), err
}

// Delete takes name of the teardown and deletes it. Returns an error if one occurs.
func (c *FakeTeardowns) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(teardownsResource, c.ns, name), &lattice_v1.Teardown{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTeardowns) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(teardownsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &lattice_v1.TeardownList{})
	return err
}

// Patch applies the patch and returns the patched teardown.
func (c *FakeTeardowns) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *lattice_v1.Teardown, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(teardownsResource, c.ns, name, data, subresources...), &lattice_v1.Teardown{})

	if obj == nil {
		return nil, err
	}
	return obj.(*lattice_v1.Teardown), err
}