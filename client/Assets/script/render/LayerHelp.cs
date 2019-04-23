

using System;
using UnityEngine;
using UnityEngine.EventSystems;
using UnityEngine.UI;

class LayerHelp : Layer
{
    public LayerHelp() : base("ui", "help")
    {
    }

    protected override bool init()
    {

        return true;
    }

    public override void uninit()
    {
        
    }

    public override void update(float delta)
    {

    }
}
