// generate by excelreader

using System.Collections.Generic;
using System.IO;
using ProtoBuf;


[ProtoContract]
public class Chapter {
	[ProtoMember(1)]
	public int Id; // ID
	[ProtoMember(2)]
	public string Name; // 名字
	[ProtoMember(3)]
	public int Gold; // 挂机金币收益
	[ProtoMember(4)]
	public int Power; // 通关战力
	
}
[ProtoContract]
public class ChapterArray {
    [ProtoMember(1)]
	public List<int> Keys;
	[ProtoMember(2)]
   	public List<Chapter> Items;
}

[ProtoContract]
public class Equip {
	[ProtoMember(1)]
	public int Id; // ID
	[ProtoMember(2)]
	public string Name; // 名字
	[ProtoMember(3)]
	public int Atk; // 攻击力
	[ProtoMember(4)]
	public int Hp; // 生命
	[ProtoMember(5)]
	public int Atkup; // 攻击成长
	[ProtoMember(6)]
	public int Hpup; // 生命成长
	
}
[ProtoContract]
public class EquipArray {
    [ProtoMember(1)]
	public List<int> Keys;
	[ProtoMember(2)]
   	public List<Equip> Items;
}

[ProtoContract]
public class Hero {
	[ProtoMember(1)]
	public int Id; // ID
	[ProtoMember(2)]
	public string Name; // 名字
	[ProtoMember(3)]
	public int Atk; // 攻击力
	[ProtoMember(4)]
	public int Hp; // 生命
	[ProtoMember(5)]
	public int Atkup; // 攻击成长
	[ProtoMember(6)]
	public int Hpup; // 生命成长
	
}
[ProtoContract]
public class HeroArray {
    [ProtoMember(1)]
	public List<int> Keys;
	[ProtoMember(2)]
   	public List<Hero> Items;
}


public class TableMgr {
    public static TableMgr ins;
    public Dictionary<int, Chapter> tableChapter;
    public Dictionary<int, Equip> tableEquip;
    public Dictionary<int, Hero> tableHero;
    

    public void loadChapter() {
        MemoryStream stream = new MemoryStream();
        byte[] data = DataLoader.Load("Chapter");
        stream.SetLength(0);
        stream.Write(data, 0, data.Length);
        stream.Seek(0, SeekOrigin.Begin);
        ChapterArray itemArray = Serializer.Deserialize<ChapterArray>(stream);
        this.tableChapter = new Dictionary<int,Chapter>();
        for (int i = 0; i < itemArray.Keys.Count; i++) {
            this.tableChapter[itemArray.Keys[i]] = itemArray.Items[i];
        }
    }
    public void loadEquip() {
        MemoryStream stream = new MemoryStream();
        byte[] data = DataLoader.Load("Equip");
        stream.SetLength(0);
        stream.Write(data, 0, data.Length);
        stream.Seek(0, SeekOrigin.Begin);
        EquipArray itemArray = Serializer.Deserialize<EquipArray>(stream);
        this.tableEquip = new Dictionary<int,Equip>();
        for (int i = 0; i < itemArray.Keys.Count; i++) {
            this.tableEquip[itemArray.Keys[i]] = itemArray.Items[i];
        }
    }
    public void loadHero() {
        MemoryStream stream = new MemoryStream();
        byte[] data = DataLoader.Load("Hero");
        stream.SetLength(0);
        stream.Write(data, 0, data.Length);
        stream.Seek(0, SeekOrigin.Begin);
        HeroArray itemArray = Serializer.Deserialize<HeroArray>(stream);
        this.tableHero = new Dictionary<int,Hero>();
        for (int i = 0; i < itemArray.Keys.Count; i++) {
            this.tableHero[itemArray.Keys[i]] = itemArray.Items[i];
        }
    }
    
    public void Load() {
        this.loadChapter();
        this.loadEquip();
        this.loadHero();
        
    }
}



