int str_length(char s[]){
    int count;
    for(count =0; s[count] != '\0'; count++);
    return count;
}

bool attendanceAward(char s[], int* Absent, int* Late){
    if(*s == '\0')
    {    return true;
    }
    
    if(Absent == 0 || Late == 0)
    {    return false;
    }
    if(*s =='P'){
        *Late = 3;
        if(!attendanceAward(s++, &Absent, &Late))
            return false;
    }
    else if(*s =='A'){
        *Late = 3;
        *Absent--;
        if(!attendanceAward(s++, &Absent, &Late))
            return false;
    }
    else if(*s =='P'){
        *Late--;
        if(!attendanceAward(s++, &Absent, &Late))
            return false;
    }
    return true;

}

bool checkRecord(char * s){
    int absent =0;
    int late = 0;
    while(*s != '\0'){
        if(*s == 'A'){
            absent++;
            late = 0;
            if(absent>1){
                return false;
            }
        }
        else if(*s == 'L'){
            late++;
            if(late>2){
                return false;
            }
        }
        else{
            late = 0;
        }
        s++;
    }
    return true;

}