# Go Tech Assignment (fail)

## General Approach

My initial take was to create an active-record style client library due to the HATEOAS-styled API
but thought it would require too much time since I was (am) a newfound Gopher.

I tried to externalize the returning HATEOAS structures (links & wrappers) into the `common` module
but in hindsight this might not be the Go way.

In theory an gRPC, MQ or any other request-response based client implementation should be able to
implement the `common.Operation` interface and leave the `accounts` module intact.

## Readings

Too many to list here

https://github.com/golang/go/wiki/CodeReviewComments
https://hassansin.github.io/Unit-Testing-http-client-in-Go
https://medium.com/@cep21/go-client-library-best-practices-83d877d604ca
https://medium.com/@cep21/package-proliferation-mistakes-in-go-493664cde6b9
https://github.com/form3tech-oss/go-form3
https://www.wolfe.id.au/2020/03/10/how-do-i-structure-my-go-project/
https://dev.to/plutov/writing-rest-api-client-in-go-3fkg

## Takeaway

- meaningful defaults, eg. API default URL
- io should be cancelable/timeout
- use underlyingTransport http.RoundTripper for testability

## TODO

- refactor tests using a test framework to make the code more maintainable in a TDD manner
- hide internal structs & functions from client pkg
- setup IT containerization
- ...

# Form3 Take Home Exercise

## Instructions
The goal of this exercise is to write a client library in Go to access our fake account API, which is provided as a Docker
container in the file `docker-compose.yaml` of this repository. Please refer to the
[Form3 documentation](http://api-docs.form3.tech/api.html#organisation-accounts) for information on how to interact with the API.

If you encounter any problems running the fake account API we would encourage you to do some debugging first,
before reaching out for help.

### The solution is expected to
- Be written in Go
- Contain documentation of your technical decisions
- Implement the `Create`, `Fetch`, `List` and `Delete` operations on the `accounts` resource. Note that filtering of the List operation is not required, but you should support paging
- Be well tested to the level you would expect in a commercial environment. Make sure your tests are easy to read.

#### Docker-compose
 - Add your solution to the provided docker-compose file
 - We should be able to run `docker-compose up` and see your tests run against the provided account API service 

### Please don't
- Use a code generator to write the client library
- Use (copy or otherwise) code from any third party without attribution to complete the exercise, as this will result in the test being rejected
- Use a library for your client (e.g: go-resty). Only test libraries are allowed.
- Implement an authentication scheme
- Implement support for the fields `data.attributes.private_identification`, `data.attributes.organisation_identification`
  and `data.relationships`, as they are omitted in the provided fake account API implementation
  
## How to submit your exercise
- Include your name in the README. If you are new to Go, please also mention this in the README so that we can consider this when reviewing your exercise
- Create a private [GitHub](https://help.github.com/en/articles/create-a-repo) repository, copy the `docker-compose` from this repository
- [Invite](https://help.github.com/en/articles/inviting-collaborators-to-a-personal-repository) @form3tech-interviewer-1 to your private repo
- Let us know you've completed the exercise using the link provided at the bottom of the email from our recruitment team

## License
Copyright 2019-2021 Form3 Financial Cloud

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.
