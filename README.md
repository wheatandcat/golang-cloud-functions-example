# golang-cloud-functions-example

## init

```
dep ensur
```

create Fire base service account.
this file save「serviceAccountKey.json」.

```
$ ls serviceAccountKey.json
serviceAccountKey.json
```

## deploy

```
gcloud functions deploy SaveItem --runtime go111 --trigger-http
```

```
gcloud functions deploy GetItem --runtime go111 --trigger-http
```
