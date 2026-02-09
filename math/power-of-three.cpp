class Solution {
public:
   bool powerOfThree(double n){
       if(n==0){
           return false;
       }
       if(n ==1){
           return true;
       }
       return powerOfThree(n/3);
   }
    
    bool isPowerOfThree(int n) {
        return powerOfThree(double(n)) and n>0;
        
        
    }
};