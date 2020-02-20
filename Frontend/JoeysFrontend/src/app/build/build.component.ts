import { Component, OnInit } from '@angular/core';

import {Observable} from 'rxjs';
import {map, startWith} from 'rxjs/operators';
import { HttpClient } from '@angular/common/http';


@Component({
  selector: 'app-build',
  templateUrl: './build.component.html',
  styleUrls: ['./build.component.css']
})
export class BuildComponent implements OnInit {
  options: string[] = ['One', 'Two', 'Three'];
  filteredOptions: Observable<string[]>;

  constructor(private http: HttpClient) { }

  getAllCommits(){
    this.http.get('http://192.168.64.3:30006/api/v1/commit/',)
    .pipe(map(responseData=>{
      const commitArray =[];
      for (const key in responseData){
        commitArray.push({ ...responseData[key] })
      }
      return commitArray
    }))
    .subscribe(responseData=>{
      console.log(responseData);
    });
  }

  ngOnInit(): void {
    this.getAllCommits()
  }

}
