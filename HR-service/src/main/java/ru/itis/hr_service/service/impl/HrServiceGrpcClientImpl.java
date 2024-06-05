package ru.itis.hr_service.service.impl;

import lombok.RequiredArgsConstructor;
import net.devh.boot.grpc.client.inject.GrpcClient;
import org.springframework.stereotype.Service;
import ru.itis.hr_service.dto.Cv;
import ru.itis.hr_service.dto.Job;
import ru.itis.hr_service.dto.JobCvPair;
import ru.itis.hr_service.grpc.pb.cv.CVServiceGrpc;
import ru.itis.hr_service.grpc.pb.job.GetAllRequest;
import ru.itis.hr_service.grpc.pb.job.GetAllResponse;
import ru.itis.hr_service.grpc.pb.job.JobServiceGrpc;
import ru.itis.hr_service.service.HrService;

import java.util.ArrayList;
import java.util.List;

@Service
@RequiredArgsConstructor
public class HrServiceGrpcClientImpl implements HrService {

    @GrpcClient("job-service")
    private JobServiceGrpc.JobServiceBlockingStub jobService;

    @GrpcClient("cv-service")
    private CVServiceGrpc.CVServiceBlockingStub cvService;

    @Override
    public List<JobCvPair> getPairs() {
        List<JobCvPair> pairs = new ArrayList<>();

        GetAllResponse jobsResponse = jobService.getAll(GetAllRequest.newBuilder().build());

        List<Job> jobs = jobsResponse.getJobsList().stream().map(j -> Job.builder()
                .name(j.getName())
                .experience(j.getExperience())
                .skills(j.getSkillsList())
                .description(j.getDescription())
                .city(j.getCity())
                .build()).toList();

        var cvsResponse = cvService.getAll(ru.itis.hr_service.grpc.pb.cv.GetAllRequest.newBuilder().build());

        List<Cv> cvs = cvsResponse.getCvsList().stream().map(c -> Cv.builder()
                .firstName(c.getFirstName())
                .lastName(c.getLastName())
                .skills(c.getSkillsList())
                .experience(c.getExperience())
                .city(c.getCity())
                .build()).toList();

        for (Job job : jobs) {
            for (Cv cv : cvs) {
                boolean enoughExperience = cv.getExperience() >= job.getExperience();
                boolean enoughSkills = true;
                if (enoughExperience) {
                    for (String skill : job.getSkills()) {
                        if (!cv.getSkills().contains(skill)) {
                            enoughSkills = false;
                            break;
                        }
                    }
                    if (enoughSkills) {
                        pairs.add(JobCvPair.builder()
                                .cv(cv)
                                .job(job)
                                .build());
                    }
                }
            }
        }

        return pairs;
    }
}

