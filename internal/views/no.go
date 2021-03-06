package views

import (
	"fmt"

	"github.com/derailed/k9s/internal/k8s"
	"github.com/derailed/k9s/internal/resource"
	"github.com/gdamore/tcell"
)

type nodeView struct {
	*resourceView
}

func newNodeView(t string, app *appView, list resource.List) resourceViewer {
	v := nodeView{newResourceView(t, app, list).(*resourceView)}
	{
		v.extraActionsFn = v.extraActions
		v.switchPage("no")
	}

	return &v
}

func (v *nodeView) extraActions(aa keyActions) {
	aa[KeyShiftC] = newKeyAction("Sort CPU", v.sortColCmd(7, false), true)
	aa[KeyShiftM] = newKeyAction("Sort MEM", v.sortColCmd(8, false), true)
	aa[tcell.KeyEnter] = newKeyAction("View Pods", v.showPodsCmd, true)
}

func (v *nodeView) sortColCmd(col int, asc bool) func(evt *tcell.EventKey) *tcell.EventKey {
	return func(evt *tcell.EventKey) *tcell.EventKey {
		t := v.getTV()
		t.sortCol.index, t.sortCol.asc = t.nameColIndex()+col, asc
		t.refresh()

		return nil
	}
}

func (v *nodeView) showPodsCmd(evt *tcell.EventKey) *tcell.EventKey {
	if !v.rowSelected() {
		return evt
	}

	showPods(v.app, "", "Node", v.selectedItem, "", "spec.nodeName="+v.selectedItem, v.backCmd)

	return nil
}

func (v *nodeView) backCmd(evt *tcell.EventKey) *tcell.EventKey {
	v.app.inject(v)

	return nil
}

func showPods(app *appView, ns, res, selected, labelSel, fieldSel string, b actionHandler) {
	mx := k8s.NewMetricsServer(app.conn())
	list := resource.NewPodList(app.conn(), mx, ns)

	list.SetLabelSelector(labelSel)
	list.SetFieldSelector(fieldSel)

	title := fmt.Sprintf("%s:%s Pods", res, selected)
	pv := newPodView(title, app, list)
	pv.setExtraActionsFn(func(aa keyActions) {
		aa[tcell.KeyEsc] = newKeyAction("Back", b, true)
	})
	app.inject(pv)
}
