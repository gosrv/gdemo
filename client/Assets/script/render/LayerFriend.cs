

using System;
using UnityEngine;
using UnityEngine.EventSystems;
using UnityEngine.UI;

class LayerFriend : Layer
{
    private ModuleFriend mModuleFriend = null;
    private GameObject[] mObjFriendItem = null;
    private GameObject mFriendPrefeb = null;
    private GameObject mFriendPane = null;

    public LayerFriend() : base("friend")
    {
    }

    protected override bool init()
    {
        mFriendPane = UT.getChild(rootObject(), "friend/Viewport/Content");
        mFriendPrefeb = UT.getChild(rootObject(), "friend/Viewport/Content/item");
        mFriendPrefeb.SetActive(false);

        BtnEventListener listener = null;
        listener = UT.addComponentTo<BtnEventListener>(rootObject(), "add/Button");
        listener.setClickCallback(this.OnBtnClickFriendAdd);

        return true;
    }


    public override void uninit()
    {
        
    }

    public override void update(float delta)
    {
    }

    
    public void OnBtnClickFriendAdd(PointerEventData eventData, object data)
    {

    }

    public void OnBtnClickGiveMobility(PointerEventData eventData, object data)
    {

    }

    public void OnBtnClickFetchMobility(PointerEventData eventData, object data)
    {

    }
}
