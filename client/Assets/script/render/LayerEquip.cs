

using System;
using UnityEngine;
using UnityEngine.EventSystems;
using UnityEngine.UI;

class LayerEquip : Layer
{
    public LayerEquip() : base("ui/equip")
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
