# SSH Setup

For RetroS to work, it is necessary to ensure the SSH is properly configured.

## Remote Machine

### Installing SSH server

```bash
sudo apt-get install openssh-server
```

### Running SSH server

```bash
sudo systemctl start ssh
```

### Create SSH Key Pair

```bash
ssh-keygen
```

Follow the prompts to generate the keys. By default this will create the
key pair in the `~/.ssh` directory.

### Authorizing the Public Key

Copy the public key on your machine to `.ssh/authorized_keys` on the remote machine.

```bash
cat ~/.ssh/id_rsa.pub >> ~/.ssh/authorized_keys
```

This will allow your local machine to connect without a password.

### Ensure Correct Permissions

```bash
chmod 700 ~/.ssh
chmod 600 ~/.ssh/authorized_keys
```

## Local Machine

### Copy Local Public Key to Remote Machine

```bash
ssh-copy-id pi@remote-ip-address
```
