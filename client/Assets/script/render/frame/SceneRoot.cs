

using System;
using UnityEngine;

public class SceneRoot : Layer
{
    public static SceneRoot instance = new SceneRoot();
    private Layer mRootLayer = null;

    public SceneRoot() : base("", "")
    {
    }

    public void setRoot(GameObject obj)
    {
        mRoot = obj;
    }

    protected override bool init()
    {
        return true;
    }

    public bool doInit()
    {
        Layer login = new LayerLogin();
        switchLayer(login);

        return true;
    }

    public void switchLayer(Layer layer)
    {
        if(mRootLayer != null)
        {
            removeChild(mRootLayer);
        }

        mRootLayer = layer;
        if (layer != null && layer.rootObject() != null)
        {
            addChild(layer);
        }
        
    }

    public override void uninit()
    {
        rootObject().transform.DetachChildren();
    }

    public override void update(float delta)
    {
        if (mRootLayer != null)
        {
            mRootLayer.doUpdate(delta);
        }
    }
}