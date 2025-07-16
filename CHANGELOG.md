# [0.1.0](https://github.com/andresperezl/discordgamesdk-go/compare/v0.0.2...v0.1.0) (2025-07-16)


### Bug Fixes

* **core:** add core field to LobbyClient to match usage in client.go and fix build error ([3932cfd](https://github.com/andresperezl/discordgamesdk-go/commit/3932cfd34323bc559fbdda8d18135b49a866762e))
* create changelog ([6679788](https://github.com/andresperezl/discordgamesdk-go/commit/6679788efe913405b74ad318e21dcd9747e8da05))
* **lobby:** remove misplaced imports and resolve C type issues in lobby metadata key helpers ([fa51da5](https://github.com/andresperezl/discordgamesdk-go/commit/fa51da5931691c9e5315e64d5776fe4ec13a091e))
* **lobby:** use cgo.Handle for safe callback passing in lobby creation example and core ([9c6ef4f](https://github.com/andresperezl/discordgamesdk-go/commit/9c6ef4f3c60122efaf5bb9d6ba38a42232816222))


### Features

* **activity:** add context-aware helpers for ClearActivity, SendRequestReply, SendInvite, AcceptInvite ([f75c445](https://github.com/andresperezl/discordgamesdk-go/commit/f75c4457ac13abf80b9f57cf7180e013ebe25955))
* **examples:** add comprehensive Go-friendly example demonstrating user and activity features ([795b9fe](https://github.com/andresperezl/discordgamesdk-go/commit/795b9fedd37f93e3b79bcb28eb4bfa1c88498fb2))
* **examples:** add new Go-friendly basic example for Discord SDK initialization ([e052627](https://github.com/andresperezl/discordgamesdk-go/commit/e05262746ac74a0fff91bc7383e9b08f2ae55bb4))
* **lobby:** expand LobbyEventsChannel to support all lobby event streams and add SetLobbyEvents to core.Core ([32e040c](https://github.com/andresperezl/discordgamesdk-go/commit/32e040c6f3df111baebbed6214714a4362fa3016))
* **overlay:** add context-aware helpers for all async overlay methods ([8c8c131](https://github.com/andresperezl/discordgamesdk-go/commit/8c8c13182b5ba565223b7604459c2375030c1c26))
* **sdk:** enhance Discord SDK download scripts with architecture detection ([5adcd16](https://github.com/andresperezl/discordgamesdk-go/commit/5adcd16e13c7820527b874e952da141b47521c06))
* **store:** add context-aware helpers for async SKU and entitlement fetches ([8600ce8](https://github.com/andresperezl/discordgamesdk-go/commit/8600ce8984919c24fe086051c0d264d61cec8e1c))
