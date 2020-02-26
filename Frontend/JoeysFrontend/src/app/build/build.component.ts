import { Component, OnInit } from '@angular/core';

import {Observable} from 'rxjs';
import {map, startWith} from 'rxjs/operators';
import { HttpClient } from '@angular/common/http';
import { commits } from './commit.model';


@Component({
  selector: 'app-build',
  templateUrl: './build.component.html',
  styleUrls: ['./build.component.css']
})
export class BuildComponent implements OnInit {
  isLoading = false;
  displayedColumns: string[] = ['Author', 'CommitId', 'CommitMessage','Build'];
  commitMessages: commits[]=[];
  options: string[] = ['One', 'Two', 'Three'];
  filteredOptions: Observable<string[]>;

  constructor(private http: HttpClient) { }

  getAllCommits(){
    //this.isLoading =true
    this.http.get<commits[]>('http://192.168.64.3:30006/api/v1/commit/',)
    .pipe(map(responseData =>{
      const commitArray: commits[]=[];
      for (const key in responseData){
       // console.log(key);
       // console.log(responseData[key]);
        commitArray.push({ ...responseData[key] })
      }
      return commitArray
    }))
    .subscribe(responseData=>{
      this.commitMessages = responseData;
      this.isLoading = false

    });
  }

  ngOnInit(): void {
    this.isLoading =true
    this.getAllCommits()
  }

}
