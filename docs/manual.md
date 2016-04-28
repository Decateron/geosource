# User Manual

This document outlines the current interactions which can be done via the website.

### Browsing Posts

To browse around posts on the website, you can drag around and zoom in and out on the map. This will by default show posts that are within the current viewing area, which can then be clicked on either in the map or the sidenav to see details about the post.

To adjust what posts are displayed and shown in the sidenav, you can click on the cog in the top left to bring up the settings section, where you can adjust flags to your personal preference.

##### Flags

Generally the flags limit posts which are shown, and are and'd together rather than or'd. So if "Visible" and "Favorites" are checked, only posts which are both visible AND favorited will be shown.

* **Visible** - If checked, only posts which are visible within the current map area are shown in the feed.
* **All** - If checked, all posts types of posts will be shown
* **Mine** - If all is unchecked and this is checked, only posts which were created by the current user will be shown
* **Favorites** - If all is unchecked and this is checked, only posts which are favorited by the current user will be shown
* **Subscriptions** - If all is unchecked and this is checked, only posts which are within channels to which the user is subscribed will be shown

### Signing In

To sign in, click the "Login" button in the top right, and then select a provider through which to sign in. Currently only Google is supported.

### Creating a Channel

To create a channel, click on the bottom of the two buttons in the bottom left of the screen. This will open a channel creation dialog, where you can specify the title of the channel, set it to either public or private (currently to no effect), and add fields by clicking on the plus button which a user can fill out when submitting a post to this channel. The labels help to give context for what sort of information should be submitted for a given field, so for instance if you had a plant channel and wanted a picture of the plant's leaves, you could add a image field with the label "Leaves". Fields can be set to required if you want to force users to enter information for that filed. Certain fields such as checkboxes and radiobuttons have additional details about the field's form which must be filled out, such as the number of checkboxes and the label for each of them.

##### Types

* **Text** - Some textual input can be provided by the user
* **Images** - Any number of images can be submitted by the user
* **Checkboxes** - Any options from a set can be set to true by the user
* **Radiobuttons** - A single option from a set can be set to true by the user

### Creating a Post

To create a post, click on the upper of the two buttons in the bottom left of the screen. This will open a post creation dialog, where you can select a channel and then fill out its associated form. Currently the save button and drafts do not do anything.

### Adding a Comment

When viewing a post in detail, comments can be seen at the bottom of the dialog. To add a comment, simply enter some text and click submit and it will appear if it was successfully added.

### Favoriting Posts

To favorite a post, click on the heart button in the sidenav. If a post is favorited, it will appear pink, otherwise, it will appear grey. To unfavorite a post, simply click the favorite button again.

### Subscribing to Channels

To subscribe to a channel, open the settings panel in the sidenav by clicking the cog in the top left. Then open the channels dropdown dialog and click on the plus button to subscribe to a channel. To unsubscribe from the channel, click on the red 'x' that appears if the channel is currently subscribed to.

### Non-working UI Elements

Currently, filtering by time, logging in through Facebook and Twitter, and opening channel settings have not been implemented, but there are still UI elements for them. Please ignore these elements as the currently serve no purpose.
