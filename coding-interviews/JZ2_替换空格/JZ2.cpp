class Solution {
public:
	void replaceSpace(char *str,int length) {
        string res,s=str;
        for(char x:s){
            if(x==' ')res+="%20";
            else res+=x;
        }
        strcpy(str,res.c_str());
	}
};