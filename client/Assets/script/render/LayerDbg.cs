

using System;
using UnityEngine;
using UnityEngine.EventSystems;
using UnityEngine.UI;

class LayerDbg : Layer
{
    private ModuleNetGameServ mModuleNetGameServ = null;
    private Text mTextInfo = null;

    public LayerDbg(string name) : base(name)
    {
    }

    protected override bool init()
    {
        mTextInfo = UT.getComponent<Text>(rootObject(), "info");
        return true;
    }

    public override void uninit()
    {
        
    }

    public override void update(float delta)
    {
        string info = string.Format("bytes send:{0}\nbytes recv:{1}", mModuleNetGameServ.getBytesSend(), mModuleNetGameServ.getBytesRecv());
        mTextInfo.text = info;
    }
}
