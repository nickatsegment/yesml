# yesml - YAML, but proper

⚠️ WIP: Expect git history rewrites. Don't import anything. Don't trust anything written here. No guarantees whatsoever. ⚠️

yesml is a strict subset of YAML focussed on being both human and machine writable without losing anything important.

It does this by:

- preserving non-semantic content like spacing and comments
- eliminating as many pointless "personal preference" stylistic choices as possible (enforcing a style)
- requiring a reference to a schema

Specifically: yesml-compliant parsers must be able to round-trip a yesml-compliant document *including stylistic content*. Most extant YAML parsers are not capable of this nor designed to do such a thing. 

# Why YAML?

Personally I prefer HCL(2). But practically, YAML is used everywhere, and HCL is not. Famously, when Github implemented actions, they switched away from HCL in the beta to only YAML in the GA. (If anyone has the details on UX studies or whatever that led to this decision, I am super keen.)

The thesis of yesml is that we can get a lot of the nice parts of HCL while still supporting the 65 gajillion applications that take YAML input.

As such, it is a goal of yesml that yesml documents be 100% parseable by common YAML parsers and consumers.
