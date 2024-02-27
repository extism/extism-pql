# PQL as an Extism Plugin

**Run the PQL compiler in 15+ programming languages, using Extism Host SDKs.**

See more: https://extism.org

## Example

Download the compiled `main.wasm` module [from Modsurfer](https://modsurfer.dylibso.com/module?hash=85bcfd5198efdbc30c0aab652c4a5de3c0bb0dfcdedf164131ab5f77329dff5c), or build it using [TinyGo](https://tinygo.org/) or Go.

```sh
tinygo build -o main.wasm -target wasi main.go
```

Next, try calling it using the [`extism` CLI](https://github.com/extism/cli): 

```sh
cat example.pql | extism call main.wasm compile --stdin --wasi

SELECT * FROM "StormEvents" WHERE ("DamageProperty" > 5000) AND (coalesce("EventType" = 'Thunderstorm Wind', FALSE)) ORDER BY "DamageProperty" DESC NULLS LAST LIMIT 3;
```

## Using it from other languages

Extism enables many languages to make calls to Wasm functions, and thus PQL could be used to generate SQL from virtually any application. See all the available Host SDKs where you can run PQL using Extism: 

https://extism.org/docs/concepts/host-sdk

