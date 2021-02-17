// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package metrics provides a stable interface to access implementation-defined
metrics exported by the Go runtime. This package is similar to existing functions
like runtime.ReadMemStats and debug.ReadGCStats, but significantly more general.

The set of metrics defined by this package may evolve as the runtime itself
evolves, and also enables variation across Go implementations, whose relevant
metric sets may not intersect.

Interface

Metrics are designated by a string key, rather than, for example, a field name in
a struct. The full list of supported metrics is always available in the slice of
Descriptions returned by All. Each Description also includes useful information
about the metric.

Thus, users of this API are encouraged to sample supported metrics defined by the
slice returned by All to remain compatible across Go versions. Of course, situations
arise where reading specific metrics is critical. For these cases, users are
encouraged to use build tags, and although metrics may be deprecated and removed,
users should consider this to be an exceptional and rare event, coinciding with a
very large change in a particular Go implementation.

Each metric key also has a "kind" that describes the format of the metric's value.
In the interest of not breaking users of this package, the "kind" for a given metric
is guaranteed not to change. If it must change, then a new metric will be introduced
with a new key and a new "kind."

Metric key format

As mentioned earlier, metric keys are strings. Their format is simple and well-defined,
designed to be both human and machine readable. It is split into two components,
separated by a colon: a rooted path and a unit. The choice to include the unit in
the key is motivated by compatibility: if a metric's unit changes, its semantics likely
did also, and a new key should be introduced.

For more details on the precise definition of the metric key's path and unit formats, see
the documentation of the Name field of the Description struct.

A note about floats

This package supports metrics whose values have a floating-point representation. In
order to improve ease-of-use, this package promises to never produce the following
classes of floating-point values: NaN, infinity.

Supported metrics

Below is the full list of supported metrics, ordered lexicographically.

	/gc/cycles/automatic:gc-cycles
		Count of completed GC cycles generated by the Go runtime.

	/gc/cycles/forced:gc-cycles
		Count of completed GC cycles forced by the application.

	/gc/cycles/total:gc-cycles
		Count of all completed GC cycles.

	/gc/heap/allocs-by-size:bytes
		Distribution of all objects allocated by approximate size.

	/gc/heap/frees-by-size:bytes
		Distribution of all objects freed by approximate size.

	/gc/heap/goal:bytes
		Heap size target for the end of the GC cycle.

	/gc/heap/objects:objects
		Number of objects, live or unswept, occupying heap memory.

	/gc/pauses:seconds
		Distribution individual GC-related stop-the-world pause latencies.

	/memory/classes/heap/free:bytes
		Memory that is completely free and eligible to be returned to
		the underlying system, but has not been. This metric is the
		runtime's estimate of free address space that is backed by
		physical memory.

	/memory/classes/heap/objects:bytes
		Memory occupied by live objects and dead objects that have
		not yet been marked free by the garbage collector.

	/memory/classes/heap/released:bytes
		Memory that is completely free and has been returned to
		the underlying system. This metric is the runtime's estimate of
		free address space that is still mapped into the process, but
		is not backed by physical memory.

	/memory/classes/heap/stacks:bytes
		Memory allocated from the heap that is reserved for stack
		space, whether or not it is currently in-use.

	/memory/classes/heap/unused:bytes
		Memory that is reserved for heap objects but is not currently
		used to hold heap objects.

	/memory/classes/metadata/mcache/free:bytes
		Memory that is reserved for runtime mcache structures, but
		not in-use.

	/memory/classes/metadata/mcache/inuse:bytes
		Memory that is occupied by runtime mcache structures that
		are currently being used.

	/memory/classes/metadata/mspan/free:bytes
		Memory that is reserved for runtime mspan structures, but
		not in-use.

	/memory/classes/metadata/mspan/inuse:bytes
		Memory that is occupied by runtime mspan structures that are
		currently being used.

	/memory/classes/metadata/other:bytes
		Memory that is reserved for or used to hold runtime
		metadata.

	/memory/classes/os-stacks:bytes
		Stack memory allocated by the underlying operating system.

	/memory/classes/other:bytes
		Memory used by execution trace buffers, structures for
		debugging the runtime, finalizer and profiler specials, and
		more.

	/memory/classes/profiling/buckets:bytes
		Memory that is used by the stack trace hash map used for
		profiling.

	/memory/classes/total:bytes
		All memory mapped by the Go runtime into the current process
		as read-write. Note that this does not include memory mapped
		by code called via cgo or via the syscall package.
		Sum of all metrics in /memory/classes.

	/sched/goroutines:goroutines
		Count of live goroutines.
*/
package metrics
