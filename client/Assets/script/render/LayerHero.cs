

using System;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.EventSystems;
using UnityEngine.UI;

class LayerHero : Layer
{
    private ModulePackHero modulePackHero;
    private GameObject content;
    private GameObject item;
    private List<GameObject> allItems = new List<GameObject>();
    public LayerHero() : base("ui/hero")
    {
    }

    protected override bool init()
    {
        content = UT.getChild(rootObject(), "ScrollView/Viewport/Content");
        item = UT.getChild(rootObject(), "ScrollView/Viewport/Content/item");
        item.SetActive(false);
        allItems.Add(item);

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
        var heros = Player.playerData.heroPack.heros;
        int idx = 0;
        foreach (proto.Hero hero in heros.Values)
        {
            GameObject objHero;
            if (idx < allItems.Count)
            {
                objHero = allItems[idx];                
            } 
            else
            {
                objHero = UT.createPrefeb(item);
                allItems.Add(objHero);
                objHero.transform.SetParent(content.transform);
                objHero.transform.localScale = Vector3.one;
            }
            idx++;

            objHero.SetActive(true);            
            Hero cfgHero = TableMgr.ins.tableHero[hero.id];
            UT.getComponent<Text>(objHero, "des").text = string.Format("{0}\nlev:{1}\nnum:{2}", cfgHero.Name, hero.level, hero.num);
        }
        for (int i=idx; i<allItems.Count; i++)
        {
            allItems[i].SetActive(false);
        }
    }
}
