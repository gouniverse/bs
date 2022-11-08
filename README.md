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

## Breadcrumbs

```
breadcrumbsPath := []bs.Breadcrumb{
	{
		Name: "Home",
		URL:  "/users",
		Icon: icons.Icon("bi-house-fill", 16, 16, "gray").ToHTML(),
	},
	{
		Name: "Profile",
		Icon: icons.Icon("bi-person-circle", 16, 16, "gray").ToHTML(),
		URL:  "/users/profile",
	},
	{
		Name: "Update details",
		URL:  "/users/profile/update",
	},
}

breadcrumbs := bs.Breadcrumbs(breadcrumbsPath).Style("background:#e9ecef;border-radius:10px;padding:10px;margin-bottom:10px;")
```

Result
<img src="Breadcrumbs.png" />

## Cards

```golang
card := bs.Card().ID("CardLeadsPreview").AddChildren([]*hb.Tag{
	bs.CardHeader().Child(hb.NewHeading5().HTML("Leads Preview")),
	bs.CardBody().Child(preview),
})
```
