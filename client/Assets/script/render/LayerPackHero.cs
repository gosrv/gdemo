

using System;
using UnityEngine;
using UnityEngine.EventSystems;
using UnityEngine.UI;

class LayerPackHero : Layer
{
    public LayerPackHero(string group, string name) : base(name)
    {
    }

    protected override bool init()
    {
        mapChild<LayerHeroList>("hero");
        mapChild<LayerHeroSell>("sell");

        return true;
    }

    public override void uninit()
    {
        
    }

    public override void update(float delta)
    {
        
    }
}
