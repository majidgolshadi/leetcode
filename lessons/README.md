# Note

Before you say "I'm done," run this checklist mentally. This separates Seniors from Staff.

1. The Empty/Null Case
    * What if nums is empty? nil?
    * Code check: Did I do `len(nums)-1` on an empty slice? That's a panic (index out of bounds if accessed or logic error).

2. The Trivial Case
    * Array with 1 element. String with 1 char.
    * Code check: Does my loop for right := 0 run? Yes.

3. The "All Same" Case
    * [2, 2, 2, 2] or "aaaaa".
    * Code check: Does my sliding window collapse correctly? Does my binary search get stuck?

4. The Limits
    * k = 0? k = len(nums)?
    * Integer overflow on addition?

5. The Output Requirement
    * Did the function ask for indices, values, or a boolean? (Common mistake: returning the value instead of the index in Two Sum).