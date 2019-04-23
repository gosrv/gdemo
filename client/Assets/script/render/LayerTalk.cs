

using System;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.EventSystems;
using UnityEngine.UI;

class LayerTalk : Layer
{
    private ModuleTalk mModuleTalk = null;
    
    public LayerTalk(string group, string name) : base(group, name)
    {
    }

    protected override bool init()
    {
        BtnEventListener listener = UT.addComponentTo<BtnEventListener>(rootObject(), "send");
        listener.setClickCallback(this.OnBtnClickTalk);

        return true;
    }

    public override void uninit()
    {
        
    }

    public override void update(float delta)
    {

    }

    public void OnBtnClickTalk(PointerEventData eventData, object data)
    {

    }
}
