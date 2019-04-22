using UnityEngine;

class DataLoader
{
    public static byte[] Load(string name)
    {
        TextAsset text = Resources.Load("table/" + name.ToLower()) as TextAsset;        
        return text.bytes;
    }
}
