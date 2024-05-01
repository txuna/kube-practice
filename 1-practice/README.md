# Pod Monitor 
지정한 deployment의 pod과 nodes정보를 기반으로 metrics정보를 파악한다. 파악된 metrics 정보를 기반으로 Pod의 Auto Scaling up and down을 처리한다.   

### 정책 

1. 평균 CPU 수치 기반 

2. 평균 메모리 수치 기반

3. 스케일 다운 속도 및 한번 감소량, 최소 레플리카수 

4. 스케일 업 속도 및 한번 증가량, 최대 레플리카 수

### 디테일
평균 CPU와 메모리 값을 기반으로 스케일 업과 스케일 다운을 진행한다. 진행 시 적용된 수치값을 기반으로 진행한다. 

스케일 다운 및 업은 Deployment의 replicas 필드를 수정한다.