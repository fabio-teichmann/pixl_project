# Widgets

## Definition

In computing, a **widget is an element of a graphical user interface** that displays information or provides a specific way for a user to interact with the operating system (OS) or an application.

Widgets include the following:

- icons;
- pull-down menus;
- buttons;
- selection boxes;
- progress indicators;
- on-off checkmarks;
- scroll bars;
- windows;
- window edges that let you resize the window;
- toggle buttons; and
- devices for displaying information and inviting, accepting and responding to user actions.

_Source: [widget](https://www.techtarget.com/whatis/definition/widget)_

## Implementation with fyne-package

In the `fyne` toolkit, in oder to create something we can click on, we need to implement the following interfaces:

```
type Widget interface {
    // Base functionality and state for all widgets
    // (size, position, etc).
    // Initialized with widget.ExtendBaseWidget(widget)
    CanvasObject
    // Creates the renderer which defines how the widget looks.
    CreateRenderer() WidgetRenderer
}
```

```
type WidgetRenderer interface {
    // Deprecated: Ignore.
    BackgroundColor() color.Color
    // Internal use: leave empty on implementation.
    Destroy()
    // Calculate the position of individual objects
    // that make up this widget.
    Layout(Size)
    // Minimum dimensions that this widget occupies.
    MinSize() Size
    // All objects that should be drawn.
    Objects() []CanvasObject
    // An update occurred and the widget needs to be redrwan.
    Refresh()
}
```
