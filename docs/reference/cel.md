# CEL Functions

Rottweiler uses the Common Expression Language (CEL) to provide a powerful and flexible way to define rules in policies. 

## Standard CEL Resources

Rottweiler supports the full standard CEL feature set. If you need to perform more advanced operations, we recommend these resources to get started:

- [CEL Overview](https://cel.dev/overview/cel-overview): The official spec and language guide.
- [CEL by Example](https://celbyexample.com/): A practical guide with examples for common use cases.

<br>

::: info Rottweiler CEL libraries
Rottweiler automatically includes common CEL libraries. You can use standard functions like `size()`, `has()`, or collection macros like `.exists()` and `.map()` out of the box in your policies.
::: 