# Bootstrap (bs)

Declaratve building of Bootstrap components. 

Its an extension of the HTML Builder package: https://github.com/gouniverse/hb

## Alerts

```golang
bs.Alert().Class("alert-info").HTML("Info text")

bs.Alert().Class("alert-danger").HTML("Danger text")

bs.Alert().Class("alert-success").HTML("Success text")

bs.Alert().Class("alert-warning").HTML("Warning text")
```

## Cards

```golang
card := bs.Card().ID("CardLeadsPreview").AddChildren([]*hb.Tag{
	bs.CardHeader().Child(hb.NewHeading5().HTML("Leads Preview")),
	bs.CardBody().Child(preview),
})
```
