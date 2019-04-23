
using System.Collections.Generic;


namespace proto
{

public class PlayerData {
            public BaseInfo baseInfo { get; private set; }
        public bool isDirtybaseInfo { get; private set; }
            public VipInfo vipInfo { get; private set; }
        public bool isDirtyvipInfo { get; private set; }
            public HeroPack heroPack { get; private set; }
        public bool isDirtyheroPack { get; private set; }
            public EquipPack equipPack { get; private set; }
        public bool isDirtyequipPack { get; private set; }
            public Guide guide { get; private set; }
        public bool isDirtyguide { get; private set; }
    

public void ClearDirty() {
        if (this.isDirtybaseInfo)
        {
            this.isDirtybaseInfo = false; 
                this.baseInfo.ClearDirty();
        }
    
        if (this.isDirtyvipInfo)
        {
            this.isDirtyvipInfo = false; 
                this.vipInfo.ClearDirty();
        }
    
        if (this.isDirtyheroPack)
        {
            this.isDirtyheroPack = false; 
                this.heroPack.ClearDirty();
        }
    
        if (this.isDirtyequipPack)
        {
            this.isDirtyequipPack = false; 
                this.equipPack.ClearDirty();
        }
    
        if (this.isDirtyguide)
        {
            this.isDirtyguide = false; 
                this.guide.ClearDirty();
        }
    
}
// read from proto    
public void FromProto(netproto.PlayerData pdata) { 
            if (pdata.baseInfo != null) {
                this.baseInfo.FromProto(pdata.baseInfo);
                this.isDirtybaseInfo = true;
            } 
            if (pdata.vipInfo != null) {
                this.vipInfo.FromProto(pdata.vipInfo);
                this.isDirtyvipInfo = true;
            } 
            if (pdata.heroPack != null) {
                this.heroPack.FromProto(pdata.heroPack);
                this.isDirtyheroPack = true;
            } 
            if (pdata.equipPack != null) {
                this.equipPack.FromProto(pdata.equipPack);
                this.isDirtyequipPack = true;
            } 
            if (pdata.guide != null) {
                this.guide.FromProto(pdata.guide);
                this.isDirtyguide = true;
            }
}
}
public class PlayerInfo {
            public long serverTime { get; private set; }
        public bool isDirtyserverTime { get; private set; }
            public string serverName { get; private set; }
        public bool isDirtyserverName { get; private set; }
    

public void ClearDirty() {
        if (this.isDirtyserverTime)
        {
            this.isDirtyserverTime = false;
        }
    
        if (this.isDirtyserverName)
        {
            this.isDirtyserverName = false;
        }
    
}
// read from proto    
public void FromProto(netproto.PlayerInfo pdata) {
            if (pdata.serverTime != null) {
                this.serverTime = ProtoUtil.OptVal(pdata.serverTime);
                this.isDirtyserverTime = true;
            }
            if (pdata.serverName != null) {
                this.serverName = ProtoUtil.OptVal(pdata.serverName);
                this.isDirtyserverName = true;
            }
}
}
public class BaseInfo {
            public long id { get; private set; }
        public bool isDirtyid { get; private set; }
            public string name { get; private set; }
        public bool isDirtyname { get; private set; }
            public int level { get; private set; }
        public bool isDirtylevel { get; private set; }
            public int exp { get; private set; }
        public bool isDirtyexp { get; private set; }
            public int gold { get; private set; }
        public bool isDirtygold { get; private set; }
            public int head { get; private set; }
        public bool isDirtyhead { get; private set; }
            public int diamond { get; private set; }
        public bool isDirtydiamond { get; private set; }
    

public void ClearDirty() {
        if (this.isDirtyid)
        {
            this.isDirtyid = false;
        }
    
        if (this.isDirtyname)
        {
            this.isDirtyname = false;
        }
    
        if (this.isDirtylevel)
        {
            this.isDirtylevel = false;
        }
    
        if (this.isDirtyexp)
        {
            this.isDirtyexp = false;
        }
    
        if (this.isDirtygold)
        {
            this.isDirtygold = false;
        }
    
        if (this.isDirtyhead)
        {
            this.isDirtyhead = false;
        }
    
        if (this.isDirtydiamond)
        {
            this.isDirtydiamond = false;
        }
    
}
// read from proto    
public void FromProto(netproto.BaseInfo pdata) {
            if (pdata.id != null) {
                this.id = ProtoUtil.OptVal(pdata.id);
                this.isDirtyid = true;
            }
            if (pdata.name != null) {
                this.name = ProtoUtil.OptVal(pdata.name);
                this.isDirtyname = true;
            }
            if (pdata.level != null) {
                this.level = ProtoUtil.OptVal(pdata.level);
                this.isDirtylevel = true;
            }
            if (pdata.exp != null) {
                this.exp = ProtoUtil.OptVal(pdata.exp);
                this.isDirtyexp = true;
            }
            if (pdata.gold != null) {
                this.gold = ProtoUtil.OptVal(pdata.gold);
                this.isDirtygold = true;
            }
            if (pdata.head != null) {
                this.head = ProtoUtil.OptVal(pdata.head);
                this.isDirtyhead = true;
            }
            if (pdata.diamond != null) {
                this.diamond = ProtoUtil.OptVal(pdata.diamond);
                this.isDirtydiamond = true;
            }
}
}
public class Guide {
            public long id { get; private set; }
        public bool isDirtyid { get; private set; }
    

public void ClearDirty() {
        if (this.isDirtyid)
        {
            this.isDirtyid = false;
        }
    
}
// read from proto    
public void FromProto(netproto.Guide pdata) {
            if (pdata.id != null) {
                this.id = ProtoUtil.OptVal(pdata.id);
                this.isDirtyid = true;
            }
}
}
public class Chapter {
            public List<int> heros = new List<int>();
        public bool isDirtyheros { get; private set; }
            public int level { get; private set; }
        public bool isDirtylevel { get; private set; }
            public int prizeGoldCheckTime { get; private set; }
        public bool isDirtyprizeGoldCheckTime { get; private set; }
            public int prizeEquipCheckTime { get; private set; }
        public bool isDirtyprizeEquipCheckTime { get; private set; }
            public List<Equip> prizeEquip = new List<Equip>();
        public bool isDirtyprizeEquip { get; private set; }
    

public void ClearDirty() {
        if (this.isDirtyheros)
        {
            this.isDirtyheros = false;
        }
    
        if (this.isDirtylevel)
        {
            this.isDirtylevel = false;
        }
    
        if (this.isDirtyprizeGoldCheckTime)
        {
            this.isDirtyprizeGoldCheckTime = false;
        }
    
        if (this.isDirtyprizeEquipCheckTime)
        {
            this.isDirtyprizeEquipCheckTime = false;
        }
    
        if (this.isDirtyprizeEquip)
        {
            this.isDirtyprizeEquip = false;
                    foreach (var v in this.prizeEquip)
                    {
                        v.ClearDirty();
                    }
        }
    
}
// read from proto    
public void FromProto(netproto.Chapter pdata) {
                for (int i=0; i<pdata.heros.Count; i++)
                {
                    var val = pdata.heros[i];
                    if (this.heros.Count <= val.key)
                    {
                        for (int j = this.heros.Count; j <= val.key; j++)
                        {
                            this.heros.Add(ProtoUtil.DefVal(ProtoUtil.OptVal(val.val)));
                        }
                    }
                    this.heros[ProtoUtil.OptVal(val.key)] = ProtoUtil.OptVal(val.val);
                }
            this.isDirtyheros = true;
            if (pdata.level != null) {
                this.level = ProtoUtil.OptVal(pdata.level);
                this.isDirtylevel = true;
            }
            if (pdata.prizeGoldCheckTime != null) {
                this.prizeGoldCheckTime = ProtoUtil.OptVal(pdata.prizeGoldCheckTime);
                this.isDirtyprizeGoldCheckTime = true;
            }
            if (pdata.prizeEquipCheckTime != null) {
                this.prizeEquipCheckTime = ProtoUtil.OptVal(pdata.prizeEquipCheckTime);
                this.isDirtyprizeEquipCheckTime = true;
            }
                for (int i=0; i<pdata.prizeEquip.Count; i++)
                {
                    var val = pdata.prizeEquip[i];
                    if (this.prizeEquip.Count <= val.key)
                    {
                        for (int j = this.prizeEquip.Count; j <= val.key; j++)
                        {
                            this.prizeEquip.Add(new Equip());
                        }
                    }
                    if (val.val == null)
                    {
                        this.prizeEquip[ProtoUtil.OptVal(val.key)] = null;
                    }
                    else
                    {
                        this.prizeEquip[ProtoUtil.OptVal(val.key)].FromProto(val.val);
                    }
                }
            this.isDirtyprizeEquip = true;
}
}
public class VipInfo {
            public int level { get; private set; }
        public bool isDirtylevel { get; private set; }
            public int exp { get; private set; }
        public bool isDirtyexp { get; private set; }
            public int totalRecharge { get; private set; }
        public bool isDirtytotalRecharge { get; private set; }
    

public void ClearDirty() {
        if (this.isDirtylevel)
        {
            this.isDirtylevel = false;
        }
    
        if (this.isDirtyexp)
        {
            this.isDirtyexp = false;
        }
    
        if (this.isDirtytotalRecharge)
        {
            this.isDirtytotalRecharge = false;
        }
    
}
// read from proto    
public void FromProto(netproto.VipInfo pdata) {
            if (pdata.level != null) {
                this.level = ProtoUtil.OptVal(pdata.level);
                this.isDirtylevel = true;
            }
            if (pdata.exp != null) {
                this.exp = ProtoUtil.OptVal(pdata.exp);
                this.isDirtyexp = true;
            }
            if (pdata.totalRecharge != null) {
                this.totalRecharge = ProtoUtil.OptVal(pdata.totalRecharge);
                this.isDirtytotalRecharge = true;
            }
}
}
public class Hero {
            public int id { get; private set; }
        public bool isDirtyid { get; private set; }
            public int level { get; private set; }
        public bool isDirtylevel { get; private set; }
            public long status { get; private set; }
        public bool isDirtystatus { get; private set; }
            public List<Equip> equips = new List<Equip>();
        public bool isDirtyequips { get; private set; }
    

public void ClearDirty() {
        if (this.isDirtyid)
        {
            this.isDirtyid = false;
        }
    
        if (this.isDirtylevel)
        {
            this.isDirtylevel = false;
        }
    
        if (this.isDirtystatus)
        {
            this.isDirtystatus = false;
        }
    
        if (this.isDirtyequips)
        {
            this.isDirtyequips = false;
                    foreach (var v in this.equips)
                    {
                        v.ClearDirty();
                    }
        }
    
}
// read from proto    
public void FromProto(netproto.Hero pdata) {
            if (pdata.id != null) {
                this.id = ProtoUtil.OptVal(pdata.id);
                this.isDirtyid = true;
            }
            if (pdata.level != null) {
                this.level = ProtoUtil.OptVal(pdata.level);
                this.isDirtylevel = true;
            }
            if (pdata.status != null) {
                this.status = ProtoUtil.OptVal(pdata.status);
                this.isDirtystatus = true;
            }
                for (int i=0; i<pdata.equips.Count; i++)
                {
                    var val = pdata.equips[i];
                    if (this.equips.Count <= val.key)
                    {
                        for (int j = this.equips.Count; j <= val.key; j++)
                        {
                            this.equips.Add(new Equip());
                        }
                    }
                    if (val.val == null)
                    {
                        this.equips[ProtoUtil.OptVal(val.key)] = null;
                    }
                    else
                    {
                        this.equips[ProtoUtil.OptVal(val.key)].FromProto(val.val);
                    }
                }
            this.isDirtyequips = true;
}
}
public class HeroPack {
            public int limit { get; private set; }
        public bool isDirtylimit { get; private set; }
            public Dictionary<int, Hero> heros = new Dictionary<int, Hero>();
        public bool isDirtyheros { get; private set; }
    

public void ClearDirty() {
        if (this.isDirtylimit)
        {
            this.isDirtylimit = false;
        }
    
        if (this.isDirtyheros)
        {
            this.isDirtyheros = false;
                    foreach (var v in this.heros.Values)
                    {
                        v.ClearDirty();
                    }
        }
    
}
// read from proto    
public void FromProto(netproto.HeroPack pdata) {
            if (pdata.limit != null) {
                this.limit = ProtoUtil.OptVal(pdata.limit);
                this.isDirtylimit = true;
            }
                for (int i=0; i<pdata.heros.Count; i++)
                {
                    var val = pdata.heros[i];
                    if (val.val == null)
                    {
                        this.heros.Remove(ProtoUtil.OptVal(val.key));
                    }
                    else
                    {
                            var oval = this.heros[ProtoUtil.OptVal(val.key)];
                            if (oval == null)
                            {
                                oval = new Hero();
                                oval.FromProto(val.val);
                                this.heros[ProtoUtil.OptVal(val.key)] = oval;
                            }
                            else
                            {
                                oval.FromProto(val.val);
                            }
                    }
                }
            this.isDirtyheros = true;
}
}
public class Equip {
            public int id { get; private set; }
        public bool isDirtyid { get; private set; }
            public int level { get; private set; }
        public bool isDirtylevel { get; private set; }
            public long status { get; private set; }
        public bool isDirtystatus { get; private set; }
            public int num { get; private set; }
        public bool isDirtynum { get; private set; }
    

public void ClearDirty() {
        if (this.isDirtyid)
        {
            this.isDirtyid = false;
        }
    
        if (this.isDirtylevel)
        {
            this.isDirtylevel = false;
        }
    
        if (this.isDirtystatus)
        {
            this.isDirtystatus = false;
        }
    
        if (this.isDirtynum)
        {
            this.isDirtynum = false;
        }
    
}
// read from proto    
public void FromProto(netproto.Equip pdata) {
            if (pdata.id != null) {
                this.id = ProtoUtil.OptVal(pdata.id);
                this.isDirtyid = true;
            }
            if (pdata.level != null) {
                this.level = ProtoUtil.OptVal(pdata.level);
                this.isDirtylevel = true;
            }
            if (pdata.status != null) {
                this.status = ProtoUtil.OptVal(pdata.status);
                this.isDirtystatus = true;
            }
            if (pdata.num != null) {
                this.num = ProtoUtil.OptVal(pdata.num);
                this.isDirtynum = true;
            }
}
}
public class EquipPack {
            public int limit { get; private set; }
        public bool isDirtylimit { get; private set; }
            public Dictionary<int, Equip> equips = new Dictionary<int, Equip>();
        public bool isDirtyequips { get; private set; }
    

public void ClearDirty() {
        if (this.isDirtylimit)
        {
            this.isDirtylimit = false;
        }
    
        if (this.isDirtyequips)
        {
            this.isDirtyequips = false;
                    foreach (var v in this.equips.Values)
                    {
                        v.ClearDirty();
                    }
        }
    
}
// read from proto    
public void FromProto(netproto.EquipPack pdata) {
            if (pdata.limit != null) {
                this.limit = ProtoUtil.OptVal(pdata.limit);
                this.isDirtylimit = true;
            }
                for (int i=0; i<pdata.equips.Count; i++)
                {
                    var val = pdata.equips[i];
                    if (val.val == null)
                    {
                        this.equips.Remove(ProtoUtil.OptVal(val.key));
                    }
                    else
                    {
                            var oval = this.equips[ProtoUtil.OptVal(val.key)];
                            if (oval == null)
                            {
                                oval = new Equip();
                                oval.FromProto(val.val);
                                this.equips[ProtoUtil.OptVal(val.key)] = oval;
                            }
                            else
                            {
                                oval.FromProto(val.val);
                            }
                    }
                }
            this.isDirtyequips = true;
}
}
public class Daily {
            public long lastUpdateTime { get; private set; }
        public bool isDirtylastUpdateTime { get; private set; }
            public bool sign { get; private set; }
        public bool isDirtysign { get; private set; }
            public int totalSign { get; private set; }
        public bool isDirtytotalSign { get; private set; }
    

public void ClearDirty() {
        if (this.isDirtylastUpdateTime)
        {
            this.isDirtylastUpdateTime = false;
        }
    
        if (this.isDirtysign)
        {
            this.isDirtysign = false;
        }
    
        if (this.isDirtytotalSign)
        {
            this.isDirtytotalSign = false;
        }
    
}
// read from proto    
public void FromProto(netproto.Daily pdata) {
            if (pdata.lastUpdateTime != null) {
                this.lastUpdateTime = ProtoUtil.OptVal(pdata.lastUpdateTime);
                this.isDirtylastUpdateTime = true;
            }
            if (pdata.sign != null) {
                this.sign = ProtoUtil.OptVal(pdata.sign);
                this.isDirtysign = true;
            }
            if (pdata.totalSign != null) {
                this.totalSign = ProtoUtil.OptVal(pdata.totalSign);
                this.isDirtytotalSign = true;
            }
}
}
public class TestList {
            public long id { get; private set; }
        public bool isDirtyid { get; private set; }
            public List<int> ListPrimitiveInt = new List<int>();
        public bool isDirtyListPrimitiveInt { get; private set; }
            public List<string> ListPrimitiveStr = new List<string>();
        public bool isDirtyListPrimitiveStr { get; private set; }
            public List<PlayerData> ListPrimitiveCom = new List<PlayerData>();
        public bool isDirtyListPrimitiveCom { get; private set; }
    

public void ClearDirty() {
        if (this.isDirtyid)
        {
            this.isDirtyid = false;
        }
    
        if (this.isDirtyListPrimitiveInt)
        {
            this.isDirtyListPrimitiveInt = false;
        }
    
        if (this.isDirtyListPrimitiveStr)
        {
            this.isDirtyListPrimitiveStr = false;
        }
    
        if (this.isDirtyListPrimitiveCom)
        {
            this.isDirtyListPrimitiveCom = false;
                    foreach (var v in this.ListPrimitiveCom)
                    {
                        v.ClearDirty();
                    }
        }
    
}
// read from proto    
public void FromProto(netproto.TestList pdata) {
            if (pdata.id != null) {
                this.id = ProtoUtil.OptVal(pdata.id);
                this.isDirtyid = true;
            }
                for (int i=0; i<pdata.ListPrimitiveInt.Count; i++)
                {
                    var val = pdata.ListPrimitiveInt[i];
                    if (this.ListPrimitiveInt.Count <= val.key)
                    {
                        for (int j = this.ListPrimitiveInt.Count; j <= val.key; j++)
                        {
                            this.ListPrimitiveInt.Add(ProtoUtil.DefVal(ProtoUtil.OptVal(val.val)));
                        }
                    }
                    this.ListPrimitiveInt[ProtoUtil.OptVal(val.key)] = ProtoUtil.OptVal(val.val);
                }
            this.isDirtyListPrimitiveInt = true;
                for (int i=0; i<pdata.ListPrimitiveStr.Count; i++)
                {
                    var val = pdata.ListPrimitiveStr[i];
                    if (this.ListPrimitiveStr.Count <= val.key)
                    {
                        for (int j = this.ListPrimitiveStr.Count; j <= val.key; j++)
                        {
                            this.ListPrimitiveStr.Add(ProtoUtil.DefVal(ProtoUtil.OptVal(val.val)));
                        }
                    }
                    this.ListPrimitiveStr[ProtoUtil.OptVal(val.key)] = ProtoUtil.OptVal(val.val);
                }
            this.isDirtyListPrimitiveStr = true;
                for (int i=0; i<pdata.ListPrimitiveCom.Count; i++)
                {
                    var val = pdata.ListPrimitiveCom[i];
                    if (this.ListPrimitiveCom.Count <= val.key)
                    {
                        for (int j = this.ListPrimitiveCom.Count; j <= val.key; j++)
                        {
                            this.ListPrimitiveCom.Add(new PlayerData());
                        }
                    }
                    if (val.val == null)
                    {
                        this.ListPrimitiveCom[ProtoUtil.OptVal(val.key)] = null;
                    }
                    else
                    {
                        this.ListPrimitiveCom[ProtoUtil.OptVal(val.key)].FromProto(val.val);
                    }
                }
            this.isDirtyListPrimitiveCom = true;
}
}
public class TestMap {
            public BaseInfo id { get; private set; }
        public bool isDirtyid { get; private set; }
            public Dictionary<int, int> MapPrimitiveIntInt = new Dictionary<int, int>();
        public bool isDirtyMapPrimitiveIntInt { get; private set; }
            public Dictionary<int, string> MapPrimitiveIntStr = new Dictionary<int, string>();
        public bool isDirtyMapPrimitiveIntStr { get; private set; }
            public Dictionary<string, int> MapPrimitiveStrInt = new Dictionary<string, int>();
        public bool isDirtyMapPrimitiveStrInt { get; private set; }
            public Dictionary<string, string> MapPrimitiveStrStr = new Dictionary<string, string>();
        public bool isDirtyMapPrimitiveStrStr { get; private set; }
            public Dictionary<int, BaseInfo> MapPrimitiveIntCom = new Dictionary<int, BaseInfo>();
        public bool isDirtyMapPrimitiveIntCom { get; private set; }
            public Dictionary<string, PlayerData> MapPrimitiveStrCom = new Dictionary<string, PlayerData>();
        public bool isDirtyMapPrimitiveStrCom { get; private set; }
    

public void ClearDirty() {
        if (this.isDirtyid)
        {
            this.isDirtyid = false; 
                this.id.ClearDirty();
        }
    
        if (this.isDirtyMapPrimitiveIntInt)
        {
            this.isDirtyMapPrimitiveIntInt = false;
        }
    
        if (this.isDirtyMapPrimitiveIntStr)
        {
            this.isDirtyMapPrimitiveIntStr = false;
        }
    
        if (this.isDirtyMapPrimitiveStrInt)
        {
            this.isDirtyMapPrimitiveStrInt = false;
        }
    
        if (this.isDirtyMapPrimitiveStrStr)
        {
            this.isDirtyMapPrimitiveStrStr = false;
        }
    
        if (this.isDirtyMapPrimitiveIntCom)
        {
            this.isDirtyMapPrimitiveIntCom = false;
                    foreach (var v in this.MapPrimitiveIntCom.Values)
                    {
                        v.ClearDirty();
                    }
        }
    
        if (this.isDirtyMapPrimitiveStrCom)
        {
            this.isDirtyMapPrimitiveStrCom = false;
                    foreach (var v in this.MapPrimitiveStrCom.Values)
                    {
                        v.ClearDirty();
                    }
        }
    
}
// read from proto    
public void FromProto(netproto.TestMap pdata) { 
            if (pdata.id != null) {
                this.id.FromProto(pdata.id);
                this.isDirtyid = true;
            }
                for (int i=0; i<pdata.MapPrimitiveIntInt.Count; i++)
                {
                    var val = pdata.MapPrimitiveIntInt[i];
                    if (val.val == null)
                    {
                        this.MapPrimitiveIntInt.Remove(ProtoUtil.OptVal(val.key));
                    }
                    else
                    {
                            this.MapPrimitiveIntInt[ProtoUtil.OptVal(val.key)] = ProtoUtil.OptVal(val.val);
                    }
                }
            this.isDirtyMapPrimitiveIntInt = true;
                for (int i=0; i<pdata.MapPrimitiveIntStr.Count; i++)
                {
                    var val = pdata.MapPrimitiveIntStr[i];
                    if (val.val == null)
                    {
                        this.MapPrimitiveIntStr.Remove(ProtoUtil.OptVal(val.key));
                    }
                    else
                    {
                            this.MapPrimitiveIntStr[ProtoUtil.OptVal(val.key)] = ProtoUtil.OptVal(val.val);
                    }
                }
            this.isDirtyMapPrimitiveIntStr = true;
                for (int i=0; i<pdata.MapPrimitiveStrInt.Count; i++)
                {
                    var val = pdata.MapPrimitiveStrInt[i];
                    if (val.val == null)
                    {
                        this.MapPrimitiveStrInt.Remove(ProtoUtil.OptVal(val.key));
                    }
                    else
                    {
                            this.MapPrimitiveStrInt[ProtoUtil.OptVal(val.key)] = ProtoUtil.OptVal(val.val);
                    }
                }
            this.isDirtyMapPrimitiveStrInt = true;
                for (int i=0; i<pdata.MapPrimitiveStrStr.Count; i++)
                {
                    var val = pdata.MapPrimitiveStrStr[i];
                    if (val.val == null)
                    {
                        this.MapPrimitiveStrStr.Remove(ProtoUtil.OptVal(val.key));
                    }
                    else
                    {
                            this.MapPrimitiveStrStr[ProtoUtil.OptVal(val.key)] = ProtoUtil.OptVal(val.val);
                    }
                }
            this.isDirtyMapPrimitiveStrStr = true;
                for (int i=0; i<pdata.MapPrimitiveIntCom.Count; i++)
                {
                    var val = pdata.MapPrimitiveIntCom[i];
                    if (val.val == null)
                    {
                        this.MapPrimitiveIntCom.Remove(ProtoUtil.OptVal(val.key));
                    }
                    else
                    {
                            var oval = this.MapPrimitiveIntCom[ProtoUtil.OptVal(val.key)];
                            if (oval == null)
                            {
                                oval = new BaseInfo();
                                oval.FromProto(val.val);
                                this.MapPrimitiveIntCom[ProtoUtil.OptVal(val.key)] = oval;
                            }
                            else
                            {
                                oval.FromProto(val.val);
                            }
                    }
                }
            this.isDirtyMapPrimitiveIntCom = true;
                for (int i=0; i<pdata.MapPrimitiveStrCom.Count; i++)
                {
                    var val = pdata.MapPrimitiveStrCom[i];
                    if (val.val == null)
                    {
                        this.MapPrimitiveStrCom.Remove(ProtoUtil.OptVal(val.key));
                    }
                    else
                    {
                            var oval = this.MapPrimitiveStrCom[ProtoUtil.OptVal(val.key)];
                            if (oval == null)
                            {
                                oval = new PlayerData();
                                oval.FromProto(val.val);
                                this.MapPrimitiveStrCom[ProtoUtil.OptVal(val.key)] = oval;
                            }
                            else
                            {
                                oval.FromProto(val.val);
                            }
                    }
                }
            this.isDirtyMapPrimitiveStrCom = true;
}
}
}
