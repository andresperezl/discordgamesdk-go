#!/usr/bin/env bash
set -e

# Disable GPG signing for this operation
git config commit.gpgsign false

git filter-branch --env-filter '
case $GIT_COMMIT in
  a7fc6478c1f3bed40378f1f9d7e9c75640368bb1)
    export GIT_AUTHOR_DATE="$(git show -s --format=%aD $GIT_COMMIT)"
    export GIT_COMMITTER_DATE="$(git show -s --format=%cD $GIT_COMMIT)"
    export GIT_AUTHOR_NAME="$(git show -s --format=%an $GIT_COMMIT)"
    export GIT_AUTHOR_EMAIL="$(git show -s --format=%ae $GIT_COMMIT)"
    export GIT_COMMITTER_NAME="$(git show -s --format=%cn $GIT_COMMIT)"
    export GIT_COMMITTER_EMAIL="$(git show -s --format=%ce $GIT_COMMIT)"
    export GIT_COMMIT_MESSAGE="feat: initial commit - Discord Game SDK Go bindings with enhanced callback handling"
    ;;
  89d433cdc503f58c0ff82a5f35821273ca0ea7fa)
    export GIT_COMMIT_MESSAGE="chore: make Discord library more idiomatic Go, add client wrappers, async patterns, builder patterns, and update examples"
    ;;
  9df896b318ed1d92ad29b1bef137674082bb7e61)
    export GIT_COMMIT_MESSAGE="chore: move non-client files to core package, update references, fix examples to use core types/constants"
    ;;
  e1453123c31fb77954769e14821987e4089f8ffc)
    export GIT_COMMIT_MESSAGE="fix: correct Windows batch commands in Makefile for DLL copying"
    ;;
  8701c86141b9bcf770c8bf575e034068352700f0)
    export GIT_COMMIT_MESSAGE="feat: enhance Discord SDK download scripts with architecture detection"
    ;;
  9f18c3c2384ff3769f515f78f69223d6e661266d)
    export GIT_COMMIT_MESSAGE="chore: update .gitignore to prevent tracking of temp files and binaries"
    ;;
  ca0fdcf8f8c8227f40eb849cb79f77874d78174f)
    export GIT_COMMIT_MESSAGE="fix: correct Windows .exe extension for example builds"
    ;;
  3cda47e2b7c60d2a61fd77c49c2fbe68caec716a)
    export GIT_COMMIT_MESSAGE="chore: move StoreManager SKU memory management to discordcgo, remove malloc/free from core"
    ;;
  b3a49ad04ae780d2f457fa8724eb08a9a3bf2ee1)
    export GIT_COMMIT_MESSAGE="chore: move StoreManager entitlement memory management to discordcgo, remove malloc/free from core"
    ;;
  5250aa7df0021414111acec5bea43b2ac39bf9ac)
    export GIT_COMMIT_MESSAGE="chore: remove unimplementable network peer/channel stub methods, document SDK limitation"
    ;;
  1ae0e289be3c0c95ec2e8e58323a3099375b86e4)
    export GIT_COMMIT_MESSAGE="chore: remove unimplementable lobby client stubs, document SDK limitations and transaction TODOs"
    ;;
  91893c6b263c3f65ec59a2cb56b6b9593e63dc5c)
    export GIT_COMMIT_MESSAGE="fix: correct type mismatch in LobbyManagerGetLobbyActivitySecret Go binding (C.DiscordLobbySecret)"
    ;;
  fe443f0a624a019dfcc3d6c75338e9f0d63bf4cb)
    export GIT_COMMIT_MESSAGE="chore: remove C imports from core, fix LobbyManagerGetLobbyGo, re-establish project structure"
    ;;
  815ae091731362918ed180fb83a7dcf015d1a2ce)
    export GIT_COMMIT_MESSAGE="fix: remove misplaced imports, resolve C type issues in lobby metadata key helpers"
    ;;
  78c7eb1142f475412734166848da2b809de01134)
    export GIT_COMMIT_MESSAGE="fix: add core field to LobbyClient to match usage in client.go, fix build error"
    ;;
  0e63eaec13224de9ca811f66170d048c6044c74d)
    export GIT_COMMIT_MESSAGE="chore: route all Discord Game SDK calls through single-threaded dispatcher for thread safety and Go idiomatic access"
    ;;
  b2b95faa29c7643ec9a8b3965c3a14297549ade2)
    export GIT_COMMIT_MESSAGE="fix: implement dispatcher re-entrancy fix for Windows thread safety"
    ;;
  bd748a033694d817bd05ad07b2dd7e1edb7e2126)
    export GIT_COMMIT_MESSAGE="feat: implement missing functions in client files and core package"
    ;;
esac
' --msg-filter 'cat <<< "$GIT_COMMIT_MESSAGE"' -- --all

echo "\nAll commit messages have been rewritten to follow semantic commit conventions." 