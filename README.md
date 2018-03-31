# Password cracking challenges

Our (AgileBits') goals in offering these challenges is to gain a better
sense of the resistance of various types of user Master Passwords to cracking if 1Password data is captured from a user's device.

## Background

Our use of Two Secret Key Derivation (2SKD) protects users from Master Password cracking attempts in the event that data is captured from our servers, but 2SKD does not offer that protection if data is captured from the user's own device. Thus the strength of user Master Passwords remains an important part of user security for 1Password.

We need to encourage users to use Master Passwords which they can 

1. People can remember
2. People can reasonably enter on their devices
3. Are sufficiently strong

The challenges are designed to help us better understand (3).

## Challenges can't be too hard (or too easy)

We would love for people to use 1Password Master Passwords that are simply too hard to crack in the event that data is captured from their local devices. But if we present cracking challenges that are too hard to win, nobody will take the challenge. Instead, we are offering what we hope are winnable challenges with sufficient prizes that many of them will be won.

Let us emphasize this point for when results come in: **These challenges are intended to be winnable. A success does not indicate any weakness in 1Password.**

This means that the passwords we present here are weaker than we recommend as 1Password Master Passwords. The prizes we offer should be worth the effort that the participants need to put in.

At the same time, we want the attempts to take some real effort so that we can get more data on that effort. In any cracking effort, there are some fixed costs of simply setting up the cracking run (preparing the data, configuring the software, etc), we want those costs to be dominated by the actual cracking.

Whether or not the difficulty of the challenges and the prizes set achieve our aims is something we will not know until we run this.

## How we help the participants

Our interest is to understand cracking efforts in terms of the strength of a test Master Password under the assumption that an attacker fully knows the details of key derivation and password generation scheme. Therefore we try to provide everything a participant will need to know to set up their systems prior to the beginning of the competition. Thus we include

1. The source for the scripts used to generate the challenge passwords
2. Sample challenges (some with "answers") published prior to the official challenge.
3. The KDF we use for these challenges is stripped of many of the idiosyncracies of the 1Password KDF that are not relevant for the difficulty of cracking locally captured data.

Individual challenges will look something like this.

```
{
    "id": "aXw39qx7a5kt",
    "hint": "3 words",
    "prf": "HMAC-SHA256",
    "rounds": 100000,
    "salt": "697c37f6ac7b6b992d12c8eab3269af6",
    "derived": "3e0f1903cc73b07a7070a661f8450d495cc99151ae67bcdf69a80d0391e7d62f"
}
```

They were generated using
code in [gentestpwds](./cmd/gentestpwds/) and [chcreator](./cmd/chcreator), the usage of which is described in [doc/README](./doc/README.md).

## Administration

To ensure fair handling of the contest itself and the award of payments, we are asking BugCrowd to administer this. This is a natural choice, as they both have the experience with delivering bounties, and have earned a reputation as a trusted party in dealing both with those offering bounties and those seeking them.

## Prizes

- For the first person or team to crack a three word password, we offer 4096 USD.

- For the second person or team to crack a different three word password, we offer 2048 USD.

- For the third person or team to crack yet a different three word password, we offer 1024 USD.

## Rules

To be written. Needs to include things like

1. No one at AgileBits or BugCrowd can win.
2. Winners must provide a write-up of what they did, with estimations of total cost to crack, guesses per second, the systems used.
3. Contact/submission address
4. Whatever else

## Schedule

Announcement of forthcoming challenge (when this and associated documents are published): TBA

Start of contest. Publication of the actual challenges: TBA
