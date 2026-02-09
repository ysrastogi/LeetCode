class Solution {
public:
    void countWords(vector<string> sentences, int& maxi, int index){
        if(index >=sentences.size()){
            return;
        }
        int n= sentences[index].length();
        int count = 0;
        for(int i= 0; i<n; i++){
            if(sentences[index][i] ==' '|| i== n-1){
                count++;
            }
        }
        maxi = max(maxi, count);
        countWords(sentences, maxi, index+1);
    }
    int mostWordsFound(vector<string>& sentences) {
        int maxi = INT_MIN;
        countWords(sentences, maxi, 0);
        return maxi;   
    }
};