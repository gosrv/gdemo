using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class SaveData : MonoBehaviour
{
    public class Data
    {
        public string account;
    }
    public static Data data;
    // Start is called before the first frame update
    void Awake()
    {        
        string val = PlayerPrefs.GetString("gdata", JsonUtility.ToJson(data));
        data = JsonUtility.FromJson<Data>(val); 
        if (data == null)
        {
            data = new Data();
        }
        StartCoroutine(save());
    }

    // Update is called once per frame
    IEnumerator save()
    {
        while (true)
        {
            yield return new WaitForSeconds(1.0f);
            PlayerPrefs.SetString("gdata", JsonUtility.ToJson(data));
        }
    }
}
