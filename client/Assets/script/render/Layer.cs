

using System.Collections.Generic;
using System.Reflection;
using UnityEngine;
using UnityEngine.UI;

public abstract class Layer
{
    protected GameObject mRoot = null;
    protected Layer mParent = null;
    private List<Layer> mChild = new List<Layer>();
    private Dictionary<string, Text> mCacheText = new Dictionary<string, Text>();

    protected Layer()
    {

    }

    private void setRoot(GameObject root)
    {
        mRoot = root;
        connectModule();
        if (!init())
        {
            SysLog.error("layer init failed.");
        }
    }

    public bool mapChild<T>(string name) where T:Layer, new()
    {
        GameObject obj = UT.getChild(rootObject(), name);
        if (obj == null)
        {
            return false;
        }
        Layer layer = new T();
        layer.setRoot(obj);
        mChild.Add(layer);
        return true;
    }

    public Layer(string group, string name)
    {
        if (group.Length > 0 && name.Length > 0)
        {
            mRoot = UT.createPrefeb(group, name);
        }

        connectModule();

        if (!init())
        {
            SysLog.error("layer " + group + ":" + name + " init failed.");
        }
    }

    private void connectModule()
    {
        /**/
        System.Type type = this.GetType();
        System.Reflection.FieldInfo[] allField = type.GetFields(
                BindingFlags.Public | BindingFlags.NonPublic | BindingFlags.Instance | BindingFlags.ExactBinding);
        foreach (System.Reflection.FieldInfo field in allField)
        {
            if (!typeof(IModule).IsAssignableFrom(field.FieldType))
            {
                continue;
            }

            field.SetValue(this, ModuleMgr.instance.getModule(field.FieldType.Name));
            if (field.GetValue(this) == null)
            {
                throw new System.Exception("module not find " + field.FieldType.Name);
            }
        }
    }

    public GameObject rootObject()
    {
        return mRoot;
    }

    public void addChild(Layer child)
    {
        Vector3 oldpos = child.rootObject().transform.localPosition;
        Quaternion oldrot = child.rootObject().transform.localRotation;
        Vector3 oldscale = child.rootObject().transform.localScale;
        RectTransform rt = child.rootObject().GetComponent<RectTransform>();
        Vector3 rtScale = new Vector3();
        if (rt != null)
        {
            rtScale = rt.localScale;
        }

        //child.rootObject().transform.parent = rootObject().transform;
        child.rootObject().transform.SetParent(rootObject().transform);

        mChild.Add(child);
        child.mParent = this;

        if (rt != null)
        {
            rt.localScale = rtScale;
        }

        child.rootObject().transform.localPosition = oldpos;
        child.rootObject().transform.localRotation = oldrot;
        child.rootObject().transform.localScale = oldscale;
    }

    public void addChild(string attach, Layer child)
    {
        Vector3 oldpos = child.rootObject().transform.localPosition;
        Quaternion oldrot = child.rootObject().transform.localRotation;
        Vector3 oldscale = child.rootObject().transform.localScale;
        RectTransform rt = child.rootObject().GetComponent<RectTransform>();
        Vector3 rtScale = new Vector3();
        if (rt != null)
        {
            rtScale = rt.localScale;
        }

        GameObject at = UT.getChild(rootObject(), attach);
        child.rootObject().transform.SetParent(at.transform);
        mChild.Add(child);
        child.mParent = this;

        if (rt != null)
        {
            rt.localScale = rtScale;
        }

        child.rootObject().transform.localPosition = oldpos;
        child.rootObject().transform.localRotation = oldrot;
        child.rootObject().transform.localScale = oldscale;
    }

    public void removeChild(Layer child)
    {
        UT.destroy(child.rootObject());
        child.uninit();
        mChild.Remove(child);
        child.mParent = null;
    }


    protected abstract bool init();
    public abstract void uninit();

    public void doUpdate(float delta)
    {
        update(delta);
        foreach (Layer child in mChild)
        {
            child.doUpdate(delta);
        }
    }

    public abstract void update(float delta);
    
    public void setText(string child, string text, string color)
    {
        Text objText = null;
        if (!mCacheText.TryGetValue(child, out objText))
        {
            objText = UT.getComponent<Text>(rootObject(), child);
            mCacheText.Add(child, objText);
        }

        objText.text = string.Format("<color=#{0}>{1}</color>", color, text);
    }

    public void setText(string child, string text)
    {
        Text objText = null;
        if (!mCacheText.TryGetValue(child, out objText))
        {
            objText = UT.getComponent<Text>(rootObject(), child);
            mCacheText.Add(child, objText);
        }

        objText.text = text;
    }

    public  void setText<T>(string child, T text)
    {
        setText(child, text.ToString());
    }

    public void setText<T>(string child, T text, string color)
    {
        setText(child, text.ToString(), color);
    }

    public void setSprite(string child, Sprite spr)
    {
        Image img = UT.getComponent<Image>(rootObject(), child);
        img.sprite = spr;
    }
}
