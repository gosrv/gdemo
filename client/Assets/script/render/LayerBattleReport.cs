

using System;
using UnityEngine;
using UnityEngine.EventSystems;
using UnityEngine.UI;

class LayerBattleReport : Layer
{
    class CharData
    {
        public Vector3 rawPos;
        public GameObject obj;
    }
    private CharData[] mLeftChar;
    private CharData[] mRightChar;
    private GameObject mObjResult;

    private int mSection;
    private float mSectionTime;
    private GameObject mSecObj;
    private Vector3 mStartPos;
    private Vector3 mTargetPos;

    public LayerBattleReport(string group, string name) : base(group, name)
    {
    }

    protected override bool init()
    {
        mLeftChar = new CharData[6];
        mRightChar = new CharData[6];
        for (int i=0; i<6; i++)
        {
            mLeftChar[i] = new CharData();
            mLeftChar[i].obj = UT.getChild(rootObject(), "char" + (i + 1));
            mLeftChar[i].rawPos = mLeftChar[i].obj.transform.localPosition;
            mLeftChar[i].obj.SetActive(false);

            mRightChar[i] = new CharData();
            mRightChar[i].obj = UT.getChild(rootObject(), "char" + (i + 7));
            mRightChar[i].rawPos = mRightChar[i].obj.transform.localPosition;
            mRightChar[i].obj.SetActive(false);
        }

        mObjResult = UT.getChild(rootObject(), "result");
        mObjResult.SetActive(false);

        BtnEventListener listener = UT.addComponentTo<BtnEventListener>(rootObject(), "quit");
        listener.setClickCallback(this.onBtnClickQuit);        

        return true;
    }

    public override void uninit()
    {
        
    }

    public override void update(float delta)
    {
        if (mSecObj != null)
        {
            mSectionTime += delta;
            if (mSectionTime < 0.5f)
            {
                mSecObj.transform.localPosition = Vector3.Lerp(mStartPos, mTargetPos, mSectionTime/0.5f);
            }
            else if (mSectionTime < 1.0f)
            {
                mSecObj.transform.localPosition = Vector3.Lerp(mTargetPos, mStartPos, (mSectionTime-0.5f)/0.5f);
            }
            else
            {
                startNextSection();
            }
        }
    }

    private void onBtnClickQuit(PointerEventData eventData, object data)
    {
        mParent.removeChild(this);
    }

    private void startNextSection()
    {
        mSection++;
        mSectionTime = 0;
        foreach (CharData cd in mLeftChar)
        {
            cd.obj.transform.localPosition = cd.rawPos;
        }
        foreach (CharData cd in mRightChar)
        {
            cd.obj.transform.localPosition = cd.rawPos;
        }

        mSecObj = null;        
    }
}
