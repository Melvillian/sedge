---
sidebar_position: 2
id: run-mev-boost-goerli
---

# Run a validator with mev-boost on Goerli

This guide shows you how to setup and run a validator using
[Teku](https://github.com/Consensys/teku/) as consensus layer, with a random
execution client, and mev-boost.

First make sure you have Sedge installed and in your PATH following the
[installation guide](quickstart/install-guide.mdx).

:::tip

If you don't have Sedge in your PATH, just open your terminal on the folder
which Sedge's executable / binary is and run `./sedge` instead of only `sedge`.

:::

Run the following command from your terminal to set up a Teku consensus and
validator nodes on Goerli with a random execution client:

```
sedge generate full-node --network goerli -c teku
```

Set up your keys running the following command from your terminal:

```
sedge keys --network goerli
```

Import the keys that you just generate in the command above using the following
command:

```
sedge import-key
```

After that, you just need to run your setup with the following command:

```
sedge run
```

The `--network` flag allow you to choose the target network for the setup. To
check out supported networks run `sedge networks`. Default network is mainnet.

The `-c/-v` flag is to select the desired consensus/validator client for the
setup. If you only use one of those flags, then the same client pair will be
used for consensus and validator nodes.

There is also a `-e` flag to select the execution client. The default behavior
is to choose a randomized client, that's why if we skip the `-e` flag this time,
a randomized execution client will be used.

mev-boost is a default setting as long as Sedge supports mev-boost for the
selected client and network. If you don't want to use mev-boost in this case,
then add the `--no-mev-boost` flag to the command. Check out the project's
[README](https://github.com/Melvillian/sedge) for more information on Sedge's
mev-boost support.
