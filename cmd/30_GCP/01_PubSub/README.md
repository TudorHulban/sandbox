## How To
check if SDk and emulator are installed and that an account is authorized:
```sh
gcloud info
```
start emulator
```sh
gcloud beta emulators pubsub start
```
check that the emulator listening port is the same as the unit test env variable set
```sh
gcloud beta emulators pubsub env-init
```


## Resources
### Videos
Emulating GCP Services for Local Application Development
```html
https://www.youtube.com/watch?v=gOtcjYixfbo
```
### Article
```html
https://cloud.google.com/sdk/gcloud/reference/beta/emulators/pubsub/
http://www.zakariaamine.com/2020-04-04/testing-gcp-pusub-golang
```