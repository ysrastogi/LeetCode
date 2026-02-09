int minimumRecolors(char * blocks, int k){
    int b =0, i =0, j= 0, w= 0;
    int len = strlen(blocks);
    int ans = INT_MAX;
   
    while((j-i)<k){
        if(blocks[j] == 'W'){
            w +=1;
        }
        else if(blocks[j] == 'B'){
            b+= 1;
        }
        j++;
    }
    ans = fmin(ans, w);
    // if(len == k){return w;}
    
     while(j<len){
        
        if(blocks[i] == 'W'){
            w-=1;
        }
        else if(blocks[i] == 'B'){
            b -= 1;
        }
        if(blocks[j] == 'W'){
            w+=1;
        }
        else if(blocks[j] == 'B'){
            b += 1;
        }
        i++;
        j++;
        ans = fmin(ans, w);
    }
    //  if(len == k){
    //         return w;
    //     }
    return ans;
}