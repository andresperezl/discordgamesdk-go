# Go Idiomatic Improvements TODO

- [x] Audit the SDK and core Go API to identify functions that could benefit from context.Context or channel-based event handling. (Completed)
- [x] Implement SetActivityWithContext in ActivityClient as a context-aware idiomatic helper for activity updates. (Completed)
- [x] Add GoDoc comments and example usage for SetActivityWithContext in the root package and examples/ directory. (Completed)
- [x] Implement a channel-based event stream for activity join requests as the next idiomatic helper. (Completed)
- [x] Document and provide an example for the channel-based event stream for activity join requests. (Completed)
- [x] Investigate and resolve the test failure (exit status 0xc0000135) for GoDoc examples. (Completed)
- [x] Implement CreateLobbyWithContext as a context-aware idiomatic helper for lobby creation, with documentation and example. (Completed)
- [x] Implement WriteWithContext as a context-aware idiomatic helper for storage writes, with documentation and example. (Completed)
- [x] Implement ReadWithContext as a context-aware idiomatic helper for storage reads, with documentation and example. (Completed)
- [x] Implement UpdateLobbyWithContext as a context-aware idiomatic helper for updating lobbies, with documentation and example. (Completed)
- [x] Implement OpenVoiceSettingsWithContext as a context-aware idiomatic helper for opening overlay voice settings, with documentation and example. (Completed)
- [x] Implement GetUserWithContext as a context-aware idiomatic helper for fetching a user by ID, with documentation and example. (Completed)
- [ ] Continue identifying and implementing more Go-idiomatic helpers or improvements (e.g., context support for other async operations, more channel-based APIs). 
- [x] Implement context-aware helpers for all async methods in OverlayClient (e.g., SetLockedWithContext, OpenActivityInviteWithContext, OpenGuildInviteWithContext, OpenVoiceSettingsWithContext). (Completed)
- [x] Implement context-aware helpers for async SKU and entitlement fetches in StoreClient (e.g., FetchSkusWithContext, FetchEntitlementsWithContext). (Completed) 
- [x] Implement context-aware helpers for ActivityClient: ClearActivityWithContext, SendRequestReplyWithContext, SendInviteWithContext, AcceptInviteWithContext. (Completed) 