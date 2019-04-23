using UnityEngine;
using System.Collections;

public class UT{
    private static MainMono mMainMono = null;

    public static void setMainMono(MainMono mainMono)
    {
        mMainMono = mainMono;
    }

    public static GameObject createPrefeb(string group, string name)
    {
        string path = group + "/" + name;        
        Object prefab = Resources.Load(path);
        Object obj = mMainMono.createPrefeb(prefab);
        return (GameObject)obj;
    }

    public static GameObject createPrefeb(GameObject obj)
    {
        return (GameObject)mMainMono.createPrefeb(obj);
    }

    public static  Object loadResource(string group, string name, System.Type type)
    {
        string path = group + "/" + name;
        return Resources.Load(path);
    }

    public static GameObject getChild(GameObject obj, string child)
    {
        if (obj == null)
        {
            return null;
        }

        Transform childTrans = obj.transform.Find(child);
        if (childTrans == null)
        {
            return null;
        }

        return childTrans.gameObject;
    }

    public static void addChild(GameObject obj, GameObject child)
    {
        Vector3 oldpos = child.transform.localPosition;
        Quaternion oldrot = child.transform.localRotation;
        Vector3 oldscale = child.transform.localScale;

        child.transform.SetParent(obj.transform);

        child.transform.localPosition = oldpos;
        child.transform.localRotation = oldrot;
        child.transform.localScale = oldscale;
    }

    public static T addComponentTo<T>(GameObject obj, string child) where T : Component
    {
        GameObject childObj = getChild(obj, child);
        if (childObj == null)
        {
            return null;
        }

        return childObj.AddComponent<T>();
    }

    public static T addComponentTo<T>(GameObject obj) where T : Component
    {
        if (obj == null)
        {
            return null;
        }

        return obj.AddComponent<T>();
    }

    public static T getComponent<T>(GameObject obj, string child) where T : Component
    {
        GameObject childObj = getChild(obj, child);
        if (childObj == null)
        {
            return null;
        }

        return childObj.GetComponent<T>();
    }

    public static T getComponent<T>(GameObject obj) where T : Component
    {
        return obj.GetComponent<T>();
    }

    public static void destroy(GameObject obj)
    {
        mMainMono.destroy(obj);
    }

    public static void destroyAllChild(GameObject obj)
    {
        Transform parent = obj.transform;
        Transform[] childs = new Transform[parent.childCount];
        for (int i=0; i<childs.Length; i++)
        {
            childs[i] = parent.GetChild(i);
        }

        foreach (Transform child in childs)
        {
            UT.destroy(child.gameObject);
        }
    }

    /*[min, max)*/
    public static int rand(int min, int max)
    {
        int randi = Random.Range(0, 1000000);
        while(randi < max)
        {
            randi += Random.Range(0, 1000000);
        }

        return min + randi % (max - min);
    }
}
