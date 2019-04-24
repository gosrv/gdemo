

using System;
using UnityEngine;
using UnityEngine.EventSystems;
using UnityEngine.UI;

class LayerChapter : Layer
{
    private ModuleChapter moduleChapter;
    private ModuleServerTime moduleServerTime;

    private Text level;
    private Text prize;

    public LayerChapter() : base("ui/chapter")
    {
    }

    protected override bool init()
    {
        level = UT.getComponent<Text>(rootObject(), "txtLevel");
        prize = UT.getComponent<Text>(rootObject(), "txtPrize");
        updateShow();

        Button btn = UT.getComponent<Button>(rootObject(), "btnGetPrize");
        btn.onClick.AddListener(() =>
        {
            moduleChapter.requestGetPrize();
        });
        btn = UT.getComponent<Button>(rootObject(), "btnNext");
        btn.onClick.AddListener(() =>
        {
            moduleChapter.requestChallenge();
        });
        return true;
    }

    public override void uninit()
    {
        
    }

    public override void update(float delta)
    {
        updateShow();
    }

    public void updateShow()
    {        
        level.text = Player.playerData.chapter.level.ToString();
        long delta = moduleServerTime.getServerTime() - Player.playerData.chapter.prizeCheckTime;
        long prizeGold = TableMgr.ins.tableChapter[Player.playerData.chapter.level].Gold * delta;
        if (prizeGold < 0)
        {
            prizeGold = 0;
        }
        prize.text = prizeGold.ToString();
    }
}
