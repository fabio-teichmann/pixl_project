# Mouse Interfaces

```
type Mousable interface {
    // Mouse Button pressed.
    MouseDown(*MouseEvent)
    // Mouse Button released.
    MouseUp(*MouseEvent)
}

type Scrollable interface {
    // Mouse scroll wheel movement.
    Scrolled(*ScrollEvent)
}

type Hoverable interface {
    // Mouse pointer enters on element
    MouseIn(*MouseEvent)
    // Mouse pointer moved over an element.
    MouseMoved(*MouseEvent)
    // Mouse pointer leaves an element.
    MouseOut()
}
```
