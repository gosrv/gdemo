

using System;
using UnityEngine;
using UnityEngine.EventSystems;
using UnityEngine.UI;

class LayerHero : Layer
{
    private ModulePackHero modulePackHero;
    private GameObject content;
    private GameObject item;

    public LayerHero() : base("ui/hero")
    {
    }

    protected override bool init()
    {
        content = UT.getChild(rootObject(), "ScrollView/Viewport/Content");
        item = UT.getChild(rootObject(), "ScrollView/Viewport/Content/item");
        item.SetActive(false);

        Button btn = UT.getComponent<Button>(rootObject(), "btnGetHero");
        btn.onClick.AddListener(() =>
        {
            modulePackHero.requestDraw(1);
        });
        btn = UT.getComponent<Button>(rootObject(), "btnGetHero10");
        btn.onClick.AddListener(() =>
        {
            modulePackHero.requestDraw(10);
        });
        updateShow();
        return true;
    }

    public override void uninit()
    {
        
    }

    public override void update(float delta)
    {
        if (Player.playerData.isDirtyheroPack)
        {
            updateShow();
        }
    }

    public void updateShow()
    {
        UT.destroyAllChild(content);
        var heros = Player.playerData.heroPack.heros;
        for (int i=0; i<heros.Count; i++)
        {
            proto.Hero hero = heros[i];
            GameObject objHero = UT.createPrefeb(item);
            objHero.SetActive(true);
            objHero.transform.SetParent(content.transform);
            Hero cfgHero = TableMgr.ins.tableHero[hero.id];
            UT.getComponent<Text>(objHero, "des").text = string.Format("{0}\n{0}\n{0}", cfgHero.Name, hero.level, hero.num);
        }
    }
}
