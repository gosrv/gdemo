using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class Config : MonoBehaviour {
    public static Config ins;
    public GameServerHost gameServerHost;

    // Use this for initialization
    void Awake () {
        ins = this;
    }
	
	// Update is called once per frame
	void Update () {

    }
}


[System.Serializable]
public class GameServerHost
{
    public string ip;
    public int port;
}