package ru.itis.hr_service.dto.response;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;
import ru.itis.hr_service.dto.JobCvPair;

import java.util.List;

@NoArgsConstructor
@AllArgsConstructor
@Builder
@Data
public class GetCandidatesResponse {
    private List<JobCvPair> pairs;
}
