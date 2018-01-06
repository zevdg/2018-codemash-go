# The encoding/json package

In this lab, we'll be modifying our service to accept a json payload as input and return a json payload as output.

1. Create an `IpsumResp` struct containing a single string field `Ipsum`.  
*Note: Capitalization is important here.  The encoding/json package cannot access fields that are unexported (lowercase).*
1. At the end of your `ipsumHandler`, remove the call to `fmt.Fprint`.
1. In it's place, create a new [json.Encoder](https://golang.org/pkg/encoding/json/#NewEncoder) initialized with the `http.ResponseWriter`.
1. Create an instance of your `IpsumResp` type, set the `Ipsum` field, and pass it to your encoder's [Encode](https://golang.org/pkg/encoding/json/#Encoder.Encode) function.  
1. If the encoding fails, write a `http.StatusInternalServerError` header, and print out an error message.
1. Test your webservice to verify that you are now returning a json payload.
1. Now create an `IpsumReq` type with two int fields: `Words`and `SentenceLength`
*Note: again, casing is important*
1. At the beginning of `ipsumHandler`, create an instance of `IpsumReq` and a [json.Decoder](https://golang.org/pkg/encoding/json/#NewDecoder) from the `Body` of [http.Request](https://golang.org/pkg/net/http/#Request)
1. Use your decoder's [decode](https://golang.org/pkg/encoding/json/#Decoder.Decode) method to populate your instance of `IpsumReq` from the request body.
1. If the decoding fails, write a `http.StatusBadRequest` header, and print out an error message.