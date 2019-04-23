using UnityEngine;
using System.Collections;

public class MainMono : MonoBehaviour {
    private static int bt = System.Environment.TickCount;    
	// Use this for initialization
	void Start () {
        UT.setMainMono(this);
        SceneRoot.instance.setRoot(this.gameObject);
        FrameWork.instance.init();
    }
	
	// Update is called once per frame
	void Update () {        
        int ct = System.Environment.TickCount;
        FrameWork.instance.tick(ct - bt);
	}

    void LateUpdate()
    {
        int ct = System.Environment.TickCount;
        FrameWork.instance.latetick(ct - bt);
        bt = ct;
    }

    public UnityEngine.Object createPrefeb(UnityEngine.Object obj)
    {        
        return Instantiate(obj);
    }

    public void destroy(GameObject obj)
    {
        Destroy(obj);
    }
}
