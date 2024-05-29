package main

/**
Structure:

URL ->
URL Parser ->
segment source generator ─┬─> segment1 -> simple mapper... ─┐                       ┌─> simple mapper...
                          ├─> segment2 -> simple mapper... ─┼─> complex mapper...  ─┼─> simple mapper...
                          └─> segment3 -> simple mapper... ─┘                       └─> simple mapper...
*/
/**
About caching:

Caching happens at http requests.
Caching middle operations is not suggested.

*/
func main() {

}
